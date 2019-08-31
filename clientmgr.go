package jccclient

import (
	"context"
	"fmt"
	"time"

	jarviscrawlercore "github.com/zhs007/jccclient/proto"
)

// ClientMgrState - ClientMgr state
type ClientMgrState int

const (
	// ClientMgrStateNormal - nil
	ClientMgrStateNormal ClientMgrState = 0
	// ClientMgrStateService - service
	ClientMgrStateService ClientMgrState = 1
	// ClientMgrStateTasks - tasks
	ClientMgrStateTasks ClientMgrState = 2
)

// ClientMgr -
type ClientMgr struct {
	cfg             *Config
	db              *DB
	Tags            map[string][]*Client
	NoTags          []*Client
	Clients         map[string]*Client
	Tasks           []*Task
	MaxTaskID       int
	State           ClientMgrState
	StopServiceChan chan int
	AddTaskChan     chan int
}

// NewClientMgr - new clientmgr
func NewClientMgr(cfg *Config) (*ClientMgr, error) {
	mgr := &ClientMgr{
		cfg:             cfg,
		Tags:            make(map[string][]*Client),
		Clients:         make(map[string]*Client),
		State:           ClientMgrStateNormal,
		StopServiceChan: make(chan int),
		AddTaskChan:     make(chan int, 16),
	}

	db, err := NewDB(cfg.DBPath, "", cfg.DBEngine)
	if err != nil {
		return nil, err
	}

	mgr.db = db
	mgr.init()

	return mgr, nil
}

// init - init
func (mgr *ClientMgr) init() {
	for _, v := range mgr.cfg.Clients {
		cc := NewClient(mgr.db, v.ServAddr, v.Token)
		cc.cfg = mgr.cfg
		cc.tags = v.Tags

		mgr.Clients[v.ServAddr] = cc

		if len(v.Tags) == 0 {
			mgr.NoTags = append(mgr.NoTags, cc)
		} else {
			for _, tv := range v.Tags {
				mgr.add2tag(tv, cc)
			}
		}
	}
}

// add2tag - add client to tag
func (mgr *ClientMgr) add2tag(tag string, c *Client) {
	_, isok := mgr.Tags[tag]
	if !isok {
		mgr.Tags[tag] = []*Client{}
	}

	mgr.Tags[tag] = append(mgr.Tags[tag], c)
}

// GetClient - get a client with hostname and tag
func (mgr *ClientMgr) GetClient(tags *Tags, hostname string) *Client {
	if tags != nil {
		lst, isok := mgr.Tags[tags.Tag]
		if isok && len(lst) > 0 {
			c := findClient(tags, lst, hostname, mgr.cfg)
			if c != nil {
				c.Hosts.Hosts[hostname].LastTime = time.Now().Unix()
			}

			return c
		}
	}

	c := findClient(nil, mgr.NoTags, hostname, mgr.cfg)
	if c != nil {
		c.Hosts.Hosts[hostname].LastTime = time.Now().Unix()
	}

	return c
}

// AddTask - add a task
func (mgr *ClientMgr) AddTask(tags *Tags, task *Task) error {
	if tags != nil && !mgr.HasTag(tags.Tag) {
		return ErrNoTag
	}

	task.taskid = mgr.newTaskID()
	task.tags = tags

	if task.AnalyzePage != nil {
		hostname, err := GetHostName(task.AnalyzePage.URL)
		if err != nil {
			return err
		}

		task.hostname = hostname
	} else if task.GeoIP != nil {
		if task.GeoIP.Platform == "" || task.GeoIP.Platform == "ipvoid" {
			task.hostname = "www.ipvoid.com"
		}
	} else if task.TechInAsia != nil {
		task.hostname = "www.techinasia.com"
	}

	mgr.Tasks = append(mgr.Tasks, task)

	if mgr.State == ClientMgrStateService {
		mgr.AddTaskChan <- task.taskid
	}

	return nil
}

// StartAllTasks - start all tasks
func (mgr *ClientMgr) StartAllTasks(ctx context.Context) error {
	if mgr.State != ClientMgrStateNormal {
		return ErrInvalidClientMgrState
	}

	if len(mgr.Tasks) <= 0 {
		return nil
	}

	mgr.State = ClientMgrStateTasks

	endchan := make(chan int, 16)

	mgr.onStartTask(ctx, endchan)

	for {
		curtaskid := <-endchan

		if mgr.nextTask(ctx, endchan, curtaskid) {
			break
		}
	}

	mgr.State = ClientMgrStateNormal

	return nil
}

// StopService - stop service
func (mgr *ClientMgr) StopService() error {
	if mgr.State != ClientMgrStateService {
		return ErrInvalidClientMgrState
	}

	mgr.StopServiceChan <- 0

	return nil
}

// onStartTask
func (mgr *ClientMgr) onStartTask(ctx context.Context, endchan chan int) {
	outputLog("debug",
		fmt.Sprintf("onStartTask"))

	for _, v := range mgr.Tasks {
		if v.running {
			continue
		}

		cc := mgr.GetClient(v.tags, v.hostname)
		if cc != nil {
			cc.Running = true
			v.running = true

			go mgr.runTask(ctx, cc, v, endchan)
		}
	}
}

// StartService - start service
func (mgr *ClientMgr) StartService(ctx context.Context) error {
	if mgr.State != ClientMgrStateNormal {
		return ErrInvalidClientMgrState
	}

	mgr.State = ClientMgrStateService

	endchan := make(chan int, 16)

	for {
		isend := false

		select {
		case curtaskid := <-endchan:
			mgr.nextTask(ctx, endchan, curtaskid)
		case <-mgr.AddTaskChan:
			mgr.onStartTask(ctx, endchan)
		case <-mgr.StopServiceChan:
			isend = true
		}

		if isend {
			break
		}
	}

	mgr.State = ClientMgrStateNormal

	return nil
}

// runTask - run a task
func (mgr *ClientMgr) runTask(ctx context.Context, client *Client, task *Task, endChan chan int) error {
	outputLog("debug",
		fmt.Sprintf("runTask client - [%v]",
			client.servAddr))

	if task.AnalyzePage != nil {
		reply, err := client.analyzePage(ctx, task.hostname, task.AnalyzePage.URL,
			&task.AnalyzePage.Viewport, &task.AnalyzePage.Options)

		mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

		return err

	} else if task.GeoIP != nil {
		reply, err := client.getGeoIP(ctx, task.hostname, task.GeoIP.IP, task.GeoIP.Platform)

		mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

		return err
	} else if task.TechInAsia != nil {
		if task.TechInAsia.Mode == jarviscrawlercore.TechInAsiaMode_TIAM_JOBLIST {
			reply, err := client.getTechInAsiaJobList(ctx, task.hostname, task.TechInAsia.JobNums, task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.TechInAsia.Mode == jarviscrawlercore.TechInAsiaMode_TIAM_JOB {
			reply, err := client.getTechInAsiaJob(ctx, task.hostname, task.TechInAsia.JobCode, task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.TechInAsia.Mode == jarviscrawlercore.TechInAsiaMode_TIAM_COMPANY {
			reply, err := client.getTechInAsiaCompany(ctx, task.hostname, task.TechInAsia.CompanyCode, task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidTechInAsiaMode
	}

	outputLog("error",
		fmt.Sprintf("runTask client - [%v] error - [ErrInvalidTask]",
			client.servAddr))

	client.Running = false

	return ErrInvalidTask
}

// HasTag - has tag
func (mgr *ClientMgr) HasTag(tag string) bool {
	_, isok := mgr.Tags[tag]
	return isok
}

// onTaskEnd - on task end
func (mgr *ClientMgr) onTaskEnd(ctx context.Context, client *Client, task *Task,
	err error, reply *jarviscrawlercore.ReplyCrawler, endChan chan int) {

	if err != nil {
		outputLog("warn",
			fmt.Sprintf("onTaskEnd client - [%v] error - [%v] RetryNums = [%v]",
				client.servAddr, err, task.RetryNums))

		if task.RetryNums > 0 {
			task.RetryNums--
			task.running = false

			client.Running = false
			endChan <- 0

			return
		}

		task.running = false
	}

	go task.Callback(ctx, task, err, reply)

	time.Sleep(time.Second * time.Duration(mgr.cfg.SleepTime))

	client.Running = false
	endChan <- task.taskid
}

// nextTask - on task end
func (mgr *ClientMgr) nextTask(ctx context.Context, endChan chan int, taskid int) bool {
	outputLog("debug",
		fmt.Sprintf("nextTask taskid - [%v]", taskid))

	if taskid > 0 {
		for i, v := range mgr.Tasks {
			if v.taskid == taskid {
				mgr.Tasks = append(mgr.Tasks[:i], mgr.Tasks[i+1:]...)

				break
			}
		}
	}

	if len(mgr.Tasks) == 0 {
		outputLog("debug",
			fmt.Sprintf("nextTask no task"))

		return true
	}

	for _, v := range mgr.Tasks {
		if v.running {
			continue
		}

		cc := mgr.GetClient(v.tags, v.hostname)
		if cc != nil {
			cc.Running = true

			go mgr.runTask(ctx, cc, v, endChan)
		}
	}

	return false
}

// newTaskID -
func (mgr *ClientMgr) newTaskID() int {
	mgr.MaxTaskID++

	return mgr.MaxTaskID
}
