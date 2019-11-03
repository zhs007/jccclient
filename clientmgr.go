package jccclient

import (
	"context"
	"fmt"
	"strings"
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
	AllClients      []*Client
	Clients         map[string]*Client
	Tasks           []*Task
	MaxTaskID       int
	State           ClientMgrState
	StopServiceChan chan int
	AddTaskChan     chan *Task
}

// NewClientMgr - new clientmgr
func NewClientMgr(cfg *Config) (*ClientMgr, error) {
	mgr := &ClientMgr{
		cfg:             cfg,
		Tags:            make(map[string][]*Client),
		Clients:         make(map[string]*Client),
		State:           ClientMgrStateNormal,
		StopServiceChan: make(chan int),
		AddTaskChan:     make(chan *Task, 16),
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
		mgr.AllClients = append(mgr.AllClients, cc)

		if len(v.Tags) == 0 {
			// mgr.NoTags = append(mgr.NoTags, cc)
		} else {
			for _, tv := range v.Tags {
				mgr.add2tag(tv, cc)
			}
		}

		mgr.add2tag("", cc)
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

// // GetClient - get a client with hostname and tag
// func (mgr *ClientMgr) onClientStartTask(c *Client, hostname string) {
// 	curt := time.Now().Unix()
// 	if curt-c.Hosts.Hosts[hostname].LastTime > int64(mgr.cfg.SleepTime) {
// 		c.Hosts.Hosts[hostname].StartTime = curt
// 	} else if c.Hosts.Hosts[hostname].StartTime <= 0 {
// 		c.Hosts.Hosts[hostname].StartTime = curt
// 	}

// 	c.Hosts.Hosts[hostname].LastTime = curt
// }

// GetClient - get a client with hostname and tag
func (mgr *ClientMgr) GetClient(tags *Tags, hostname string) *Client {
	if tags != nil {
		lst, isok := mgr.Tags[tags.Tag]
		if isok && len(lst) > 0 {
			c := findClient(tags, lst, hostname, mgr.cfg)
			if c != nil {
				// mgr.onClientStartTask(c, hostname)
				// c.Hosts.Hosts[hostname].LastTime = time.Now().Unix()
			}

			return c
		}
	}

	c := findClient(nil, mgr.AllClients, hostname, mgr.cfg)
	if c != nil {
		// mgr.onClientStartTask(c, hostname)
		// c.Hosts.Hosts[hostname].LastTime = time.Now().Unix()
	}

	return c
}

// AddTask - add a task
func (mgr *ClientMgr) AddTask(tags *Tags, task *Task) error {
	if tags != nil && !mgr.HasTag(tags.Tag) {
		return ErrNoTag
	}

	task.TaskID = mgr.newTaskID()
	task.Tags = tags

	if task.AnalyzePage != nil {
		hostname, err := GetHostName(task.AnalyzePage.URL)
		if err != nil {
			return err
		}

		task.Hostname = hostname
	} else if task.GeoIP != nil {
		if task.GeoIP.Platform == "" || task.GeoIP.Platform == "ipvoid" {
			task.Hostname = "www.ipvoid.com"
		}
	} else if task.TechInAsia != nil {
		task.Hostname = "www.techinasia.com"
	}

	if mgr.State == ClientMgrStateService {
		mgr.AddTaskChan <- task
	} else {
		mgr.Tasks = append(mgr.Tasks, task)
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
		fmt.Sprintf("onStartTask maxtaskid - [%v]", mgr.MaxTaskID))

	for _, v := range mgr.Tasks {
		if v.Running {
			continue
		}

		cc := mgr.GetClient(v.Tags, v.Hostname)
		if cc != nil {
			cc.Running = true
			v.Running = true

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
	timerTick := time.NewTimer(60 * time.Second)

	for {
		isend := false

		select {
		case curtaskid := <-endchan:
			mgr.nextTask(ctx, endchan, curtaskid)
		case curtask := <-mgr.AddTaskChan:
			mgr.Tasks = append(mgr.Tasks, curtask)

			mgr.onStartTask(ctx, endchan)
		case <-mgr.StopServiceChan:
			isend = true
		case <-timerTick.C:
			mgr.nextTask(ctx, endchan, -1)
			timerTick.Reset(60 * time.Second)
		}

		if isend {
			break
		}
	}

	timerTick.Stop()
	mgr.State = ClientMgrStateNormal

	return nil
}

// runTask - run a task
func (mgr *ClientMgr) runTask(ctx context.Context, client *Client, task *Task, endChan chan int) error {
	outputLog("debug",
		fmt.Sprintf("runTask client - [%v] taskid - [%v]",
			client.servAddr, task.TaskID))

	task.ServAddr = client.servAddr

	if task.AnalyzePage != nil {
		reply, err := client.analyzePage(ctx, task.Hostname, task.AnalyzePage.URL,
			&task.AnalyzePage.Viewport, &task.AnalyzePage.Options)

		mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

		return err

	} else if task.GeoIP != nil {
		reply, err := client.getGeoIP(ctx, task.Hostname, task.GeoIP.IP, task.GeoIP.Platform)

		mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

		return err
	} else if task.TechInAsia != nil {
		if task.TechInAsia.Mode == jarviscrawlercore.TechInAsiaMode_TIAM_JOBLIST {
			reply, err := client.getTechInAsiaJobList(ctx, task.Hostname, task.TechInAsia.JobTag,
				task.TechInAsia.JobSubTag, task.TechInAsia.JobNums, task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.TechInAsia.Mode == jarviscrawlercore.TechInAsiaMode_TIAM_JOB {
			reply, err := client.getTechInAsiaJob(ctx, task.Hostname, task.TechInAsia.JobCode, task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.TechInAsia.Mode == jarviscrawlercore.TechInAsiaMode_TIAM_COMPANY {
			reply, err := client.getTechInAsiaCompany(ctx, task.Hostname, task.TechInAsia.CompanyCode, task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.TechInAsia.Mode == jarviscrawlercore.TechInAsiaMode_TIAM_JOBTAG {
			reply, err := client.getTechInAsiaJobTagList(ctx, task.Hostname, task.TechInAsia.JobTag, task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidTechInAsiaMode
	} else if task.SteepAndCheap != nil {
		if task.SteepAndCheap.Mode == jarviscrawlercore.SteepAndCheapMode_SACM_PRODUCTS {
			reply, err := client.getSteepAndCheapProducts(ctx, task.Hostname, task.SteepAndCheap.URL,
				task.SteepAndCheap.Page, task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.SteepAndCheap.Mode == jarviscrawlercore.SteepAndCheapMode_SACM_PRODUCT {
			reply, err := client.getSteepAndCheapProduct(ctx, task.Hostname, task.SteepAndCheap.URL,
				task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidSteepAndCheapMode
	} else if task.JRJ != nil {
		if task.JRJ.Mode == jarviscrawlercore.JRJMode_JRJM_FUND {
			reply, err := client.getJRJFund(ctx, task.Hostname, task.JRJ.Code, task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.JRJ.Mode == jarviscrawlercore.JRJMode_JRJM_FUNDS {
			reply, err := client.getJRJFunds(ctx, task.Hostname,
				task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.JRJ.Mode == jarviscrawlercore.JRJMode_JRJM_FUNDMANAGER {
			reply, err := client.getJRJFundManager(ctx, task.Hostname, task.JRJ.Code,
				task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.JRJ.Mode == jarviscrawlercore.JRJMode_JRJM_FUNDVALUE {
			reply, err := client.getJRJFundValue(ctx, task.Hostname, task.JRJ.Code, task.JRJ.Year,
				task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidJRJMode
	} else if task.JD != nil {
		if task.JD.Mode == jarviscrawlercore.JDMode_JDM_ACTIVE {
			reply, err := client.getJDActive(ctx, task.Hostname, task.JD.URL,
				task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.JD.Mode == jarviscrawlercore.JDMode_JDM_PRODUCT {
			reply, err := client.getJDProduct(ctx, task.Hostname, task.JD.URL,
				task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.JD.Mode == jarviscrawlercore.JDMode_JDM_ACTIVEPAGE {
			reply, err := client.getJDActivePage(ctx, task.Hostname, task.JD.URL,
				task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidJDMode
	} else if task.Alimama != nil {
		if task.Alimama.Mode == jarviscrawlercore.AlimamaMode_ALIMMM_KEEPALIVE {
			reply, err := client.alimamaKeepalive(ctx, task.Hostname,
				task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.Alimama.Mode == jarviscrawlercore.AlimamaMode_ALIMMM_GETTOP {
			reply, err := client.alimamaGetTop(ctx, task.Hostname,
				task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.Alimama.Mode == jarviscrawlercore.AlimamaMode_ALIMMM_SEARCH {
			reply, err := client.alimamaSearch(ctx, task.Hostname, task.Alimama.Text,
				task.Timeout)

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidAlimamaMode
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
			fmt.Sprintf("onTaskEnd client - [%v] error - [%v] RetryNums = [%v] task = %v",
				client.servAddr, err, task.RetryNums, task.ToString()))

		if strings.Index(err.Error(), "Error: noretry:") != 0 ||
			strings.Index(err.Error(), "noretry:") != 0 {

			if task.RetryNums > 0 {
				task.RetryNums--

				// time.Sleep(time.Second * time.Duration(mgr.cfg.SleepTime))

				task.Running = false
				client.Running = false
				endChan <- 0

				return
			}

			task.Fail = true
			// task.running = false
		}
	}

	go task.Callback(ctx, task, err, reply)

	// time.Sleep(time.Second * time.Duration(mgr.cfg.SleepTime))

	task.Running = false
	client.Running = false
	endChan <- task.TaskID
}

// nextTask - on task end
func (mgr *ClientMgr) nextTask(ctx context.Context, endChan chan int, taskid int) bool {
	outputLog("debug",
		fmt.Sprintf("nextTask taskid - [%v]", taskid))

	if taskid > 0 {
		for i, v := range mgr.Tasks {
			if v.TaskID == taskid {
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

	failnums := 0
	for _, v := range mgr.Tasks {
		if v.Running {
			continue
		}

		if v.Fail {
			failnums++
		}

		cc := mgr.GetClient(v.Tags, v.Hostname)
		if cc != nil {
			v.Running = true
			cc.Running = true

			go mgr.runTask(ctx, cc, v, endChan)
		}
	}

	if failnums == len(mgr.Tasks) {
		outputLog("debug",
			fmt.Sprintf("nextTask no task fail - [%v]", failnums))

		return true
	}

	return false
}

// newTaskID -
func (mgr *ClientMgr) newTaskID() int {
	mgr.MaxTaskID++

	return mgr.MaxTaskID
}
