package jccclient

import (
	"context"
	"time"

	jarviscrawlercore "github.com/zhs007/jccclient/proto"
)

// ClientMgr -
type ClientMgr struct {
	cfg       *Config
	db        *DB
	Tags      map[string][]*Client
	NoTags    []*Client
	Clients   map[string]*Client
	Tasks     []*Task
	MaxTaskID int
}

// NewClientMgr - new clientmgr
func NewClientMgr(cfg *Config) (*ClientMgr, error) {
	mgr := &ClientMgr{
		cfg:     cfg,
		Tags:    make(map[string][]*Client),
		Clients: make(map[string]*Client),
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
func (mgr *ClientMgr) GetClient(tag string, hostname string) *Client {
	if tag != "" {
		lst, isok := mgr.Tags[tag]
		if isok && len(lst) > 0 {
			c := findClient(lst, hostname, mgr.cfg)
			if c != nil {
				c.Hosts.Hosts[hostname].LastTime = time.Now().Unix()
			}

			return c
		}
	}

	c := findClient(mgr.NoTags, hostname, mgr.cfg)
	if c != nil {
		c.Hosts.Hosts[hostname].LastTime = time.Now().Unix()
	}

	return c
}

// AddTask - add a task
func (mgr *ClientMgr) AddTask(tag string, task *Task) error {
	if !mgr.HasTag(tag) {
		return ErrNoTag
	}

	task.taskid = mgr.newTaskID()
	task.tag = tag

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
	}

	mgr.Tasks = append(mgr.Tasks, task)

	return nil
}

// Start - start all tasks
func (mgr *ClientMgr) Start(ctx context.Context) error {
	if len(mgr.Tasks) <= 0 {
		return nil
	}

	endchan := make(chan int, 16)

	for _, v := range mgr.Tasks {
		cc := mgr.GetClient(v.tag, v.hostname)
		if cc != nil {
			cc.Running = true

			go mgr.runTask(ctx, cc, v, endchan)
		}
	}

	for {
		curtaskid := <-endchan

		if mgr.nextTask(ctx, endchan, curtaskid) {
			break
		}
	}

	return nil
}

// runTask - run a task
func (mgr *ClientMgr) runTask(ctx context.Context, client *Client, task *Task, endChan chan int) error {
	if task.AnalyzePage != nil {
		reply, err := client.analyzePage(ctx, task.hostname, task.AnalyzePage.URL,
			&task.AnalyzePage.Viewport, &task.AnalyzePage.Options)

		mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

		return err

	} else if task.GeoIP != nil {
		reply, err := client.getGeoIP(ctx, task.hostname, task.GeoIP.IP, task.GeoIP.Platform)

		mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

		return err
	}

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

	task.Callback(ctx, task, err, reply)

	client.Running = false
	endChan <- task.taskid
}

// nextTask - on task end
func (mgr *ClientMgr) nextTask(ctx context.Context, endChan chan int, taskid int) bool {
	for i, v := range mgr.Tasks {
		if v.taskid == taskid {
			mgr.Tasks = append(mgr.Tasks[:i], mgr.Tasks[i+1:]...)

			break
		}
	}

	if len(mgr.Tasks) == 0 {
		return true
	}

	for _, v := range mgr.Tasks {
		cc := mgr.GetClient(v.tag, v.hostname)
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
