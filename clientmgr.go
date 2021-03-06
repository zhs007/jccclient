package jccclient

import (
	"context"
	"time"

	"github.com/zhs007/ankadb"
	jarviscrawlercore "github.com/zhs007/jccclient/pb"
	"go.uber.org/zap"
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

// ClientMgr - 其它模块操作jccclient的主类
//		有service模式，在service模式下，线程安全
type ClientMgr struct {
	cfg             *Config
	db              *DB
	Tags            map[string][]*Client
	AllClients      []*Client
	Clients         map[string]*Client
	tasks           []*Task
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
func (mgr *ClientMgr) AddTask(tags *Tags, task *Task) (*Task, error) {
	if tags != nil && !mgr.HasTag(tags.Tag) {
		return nil, ErrNoTag
	}

	task.TaskID = mgr.newTaskID()
	task.Tags = tags

	if task.AnalyzePage != nil {
		hostname, err := GetHostName(task.AnalyzePage.URL)
		if err != nil {
			return nil, err
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
		mgr.tasks = append(mgr.tasks, task)
	}

	return task, nil
}

// StartAllTasks - start all tasks
func (mgr *ClientMgr) StartAllTasks(ctx context.Context) error {
	if mgr.State != ClientMgrStateNormal {
		return ErrInvalidClientMgrState
	}

	if len(mgr.tasks) <= 0 {
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
	mainLogger.Debug("onStartTask", zap.Int("MaxTaskID", mgr.MaxTaskID))

	for _, v := range mgr.tasks {
		if v.Running {
			continue
		}

		cc := mgr.GetClient(v.Tags, v.Hostname)
		if cc != nil {
			cc.Running = true
			v.Running = true

			go mgr.runTask(ctx, cc, v, endchan)
		}
		//  else {
		// 	if v.Logger != nil {
		// 		v.Logger.Warn("onStartTask: Can't find client", JSON("task", v))
		// 	}
		// }
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
			mgr.tasks = append(mgr.tasks, curtask)

			mgr.tasks = RebuildTasks(mgr.tasks)

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
	if task.Logger != nil {
		task.Logger.Debug("runTask", zap.String("servaddr", client.servAddr), JSON("task", task))
	}

	task.ServAddr = client.servAddr

	if task.AnalyzePage != nil {
		version, reply, err := client.analyzePage(ctx, task.Hostname, task.AnalyzePage.URL,
			&task.AnalyzePage.Viewport, &task.AnalyzePage.Options)

		task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

		mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

		return err

	} else if task.GeoIP != nil {
		version, reply, err := client.getGeoIP(ctx, task.Hostname, task.GeoIP.IP, task.GeoIP.Platform)

		task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

		mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

		return err
	} else if task.TechInAsia != nil {
		if task.TechInAsia.Mode == jarviscrawlercore.TechInAsiaMode_TIAM_JOBLIST {
			version, reply, err := client.getTechInAsiaJobList(ctx, task.Hostname, task.TechInAsia.JobTag,
				task.TechInAsia.JobSubTag, task.TechInAsia.JobNums, task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.TechInAsia.Mode == jarviscrawlercore.TechInAsiaMode_TIAM_JOB {
			version, reply, err := client.getTechInAsiaJob(ctx, task.Hostname, task.TechInAsia.JobCode,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.TechInAsia.Mode == jarviscrawlercore.TechInAsiaMode_TIAM_COMPANY {
			version, reply, err := client.getTechInAsiaCompany(ctx, task.Hostname, task.TechInAsia.CompanyCode,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.TechInAsia.Mode == jarviscrawlercore.TechInAsiaMode_TIAM_JOBTAG {
			version, reply, err := client.getTechInAsiaJobTagList(ctx, task.Hostname, task.TechInAsia.JobTag,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidTechInAsiaMode
	} else if task.SteepAndCheap != nil {
		if task.SteepAndCheap.Mode == jarviscrawlercore.SteepAndCheapMode_SACM_PRODUCTS {
			version, reply, err := client.getSteepAndCheapProducts(ctx, task.Hostname, task.SteepAndCheap.URL,
				task.SteepAndCheap.Page, task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.SteepAndCheap.Mode == jarviscrawlercore.SteepAndCheapMode_SACM_PRODUCT {
			version, reply, err := client.getSteepAndCheapProduct(ctx, task.Hostname, task.SteepAndCheap.URL,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidSteepAndCheapMode
	} else if task.JRJ != nil {
		if task.JRJ.Mode == jarviscrawlercore.JRJMode_JRJM_FUND {
			version, reply, err := client.getJRJFund(ctx, task.Hostname, task.JRJ.Code, task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.JRJ.Mode == jarviscrawlercore.JRJMode_JRJM_FUNDS {
			version, reply, err := client.getJRJFunds(ctx, task.Hostname,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.JRJ.Mode == jarviscrawlercore.JRJMode_JRJM_FUNDMANAGER {
			version, reply, err := client.getJRJFundManager(ctx, task.Hostname, task.JRJ.Code,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.JRJ.Mode == jarviscrawlercore.JRJMode_JRJM_FUNDVALUE {
			version, reply, err := client.getJRJFundValue(ctx, task.Hostname, task.JRJ.Code, task.JRJ.Year,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidJRJMode
	} else if task.JD != nil {
		if task.JD.Mode == jarviscrawlercore.JDMode_JDM_ACTIVE {
			version, reply, err := client.getJDActive(ctx, task.Hostname, task.JD.URL,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.JD.Mode == jarviscrawlercore.JDMode_JDM_PRODUCT {
			version, reply, err := client.getJDProduct(ctx, task.Hostname, task.JD.URL,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.JD.Mode == jarviscrawlercore.JDMode_JDM_ACTIVEPAGE {
			version, reply, err := client.getJDActivePage(ctx, task.Hostname, task.JD.URL,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidJDMode
	} else if task.Alimama != nil {
		if task.Alimama.Mode == jarviscrawlercore.AlimamaMode_ALIMMM_KEEPALIVE {
			version, reply, err := client.alimamaKeepalive(ctx, task.Hostname,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.Alimama.Mode == jarviscrawlercore.AlimamaMode_ALIMMM_GETTOP {
			version, reply, err := client.alimamaGetTop(ctx, task.Hostname,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.Alimama.Mode == jarviscrawlercore.AlimamaMode_ALIMMM_SEARCH {
			version, reply, err := client.alimamaSearch(ctx, task.Hostname, task.Alimama.Text,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.Alimama.Mode == jarviscrawlercore.AlimamaMode_ALIMMM_GETSHOP {
			version, reply, err := client.alimamaShop(ctx, task.Hostname, task.Alimama.URL,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidAlimamaMode
	} else if task.Tmall != nil {
		if task.Tmall.Mode == jarviscrawlercore.TmallMode_TMM_PRODUCT {
			version, reply, err := client.tmallProduct(ctx, task.Hostname, task.Tmall.ItemID,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.Tmall.Mode == jarviscrawlercore.TmallMode_TMM_MOBILEPRODUCT {
			version, reply, err := client.tmallMobileProduct(ctx, task.Hostname, task.Tmall.ItemID,
				task.Tmall.Device, task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidTmallMode
	} else if task.Taobao != nil {
		if task.Taobao.Mode == jarviscrawlercore.TaobaoMode_TBM_PRODUCT {
			version, reply, err := client.taobaoProduct(ctx, task.Hostname, task.Taobao.ItemID,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.Taobao.Mode == jarviscrawlercore.TaobaoMode_TBM_MOBILEPRODUCT {
			version, reply, err := client.taobaoMobileProduct(ctx, task.Hostname, task.Taobao.ItemID,
				task.Taobao.Device, task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.Taobao.Mode == jarviscrawlercore.TaobaoMode_TBM_SEARCH {
			version, reply, err := client.taobaoSearch(ctx, task.Hostname,
				task.Taobao.Text, task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidTmallMode
	} else if task.MountainSteals != nil {
		if task.MountainSteals.Mode == jarviscrawlercore.MountainStealsMode_MSM_SALE {
			version, reply, err := client.mountainstealsSale(ctx, task.Hostname, task.MountainSteals.URL,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.MountainSteals.Mode == jarviscrawlercore.MountainStealsMode_MSM_PRODUCT {
			version, reply, err := client.mountainstealsProduct(ctx, task.Hostname, task.MountainSteals.URL,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidMountainstealsMode
	} else if task.Douban != nil {
		if task.Douban.Mode == jarviscrawlercore.DoubanMode_DBM_SEARCH {
			version, reply, err := client.doubanSearch(ctx, task.Hostname, task.Douban.DoubanType,
				task.Douban.Text, task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.Douban.Mode == jarviscrawlercore.DoubanMode_DBM_BOOK {
			version, reply, err := client.doubanBook(ctx, task.Hostname, task.Douban.ID,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidTmallMode
	} else if task.ManhuaDB != nil {
		if task.ManhuaDB.Mode == jarviscrawlercore.ManhuaDBMode_MHDB_AUTHOR {
			version, reply, err := client.manhuadbAuthor(ctx, task.Hostname, task.ManhuaDB.AuthorID,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidTmallMode
	} else if task.OABT != nil {
		if task.OABT.Mode == jarviscrawlercore.OABTMode_OABTM_PAGE {
			version, reply, err := client.oabtPage(ctx, task.Hostname, task.OABT.PageIndex,
				task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidTmallMode
	} else if task.Hao6v != nil {
		if task.Hao6v.Mode == jarviscrawlercore.Hao6VMode_H6VM_NEWPAGE {
			version, reply, err := client.hao6vNewest(ctx, task.Hostname, task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.Hao6v.Mode == jarviscrawlercore.Hao6VMode_H6VM_RESPAGE {
			version, reply, err := client.hao6vRes(ctx, task.Hostname, task.Hao6v.URL, task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidHao6vMode
	} else if task.P6vdy != nil {
		if task.P6vdy.Mode == jarviscrawlercore.P6VdyMode_P6VDY_MOVIES {
			version, reply, err := client.p6vdyMovies(ctx, task.Hostname, task.P6vdy.URL, task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.P6vdy.Mode == jarviscrawlercore.P6VdyMode_P6VDY_MOVIE {
			version, reply, err := client.p6vdyMovie(ctx, task.Hostname, task.P6vdy.URL, task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidP6vdyMode
	} else if task.Investing != nil {
		if task.Investing.Mode == jarviscrawlercore.InvestingMode_INVESTINGMODE_ASSETS {
			version, reply, err := client.investingAssets(ctx, task.Hostname, task.Investing.URL, task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		} else if task.Investing.Mode == jarviscrawlercore.InvestingMode_INVESTINGMODE_HD {
			version, reply, err := client.investingHD(ctx, task.Hostname, task.Investing.URL, task.Investing.StartData, task.Investing.EndData, task.Timeout)

			task.JCCInfo.Nodes = append(task.JCCInfo.Nodes, JCCNode{Addr: client.servAddr, Version: version})

			mgr.onTaskEnd(ctx, client, task, err, reply, endChan)

			return err
		}

		return ErrInvalidInvestingMode
	}

	if task.Logger != nil {
		task.Logger.Error("runTask: ErrInvalidTask", zap.String("servaddr", client.servAddr), JSON("task", task))
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

	if err != nil {
		if task.Logger != nil {
			task.Logger.Warn("onTaskEnd: error",
				zap.Error(err),
				zap.String("servaddr", client.servAddr),
				JSON("task", task))
		}

		// if !(strings.Index(err.Error(), "Error: noretry:") == 0 ||
		// 	strings.Index(err.Error(), "noretry:") == 0) {
		if !IsNoRetryError(err) {

			if task.RetryNums > 0 {
				task.RetryNums--

				// time.Sleep(time.Second * time.Duration(mgr.cfg.SleepTime))

				task.Running = false
				client.Running = false
				endChan <- 0

				return
			}

			// task.Fail = true
			// task.running = false
		}

		task.Fail = true
	}

	go task.Callback(ctx, task, err, reply)

	// time.Sleep(time.Second * time.Duration(mgr.cfg.SleepTime))

	task.Running = false
	client.Running = false
	endChan <- task.TaskID
}

// nextTask - on task end
func (mgr *ClientMgr) nextTask(ctx context.Context, endChan chan int, taskid int) bool {
	mainLogger.Debug("nextTask", zap.Int("taskid", taskid))

	if taskid > 0 {
		for i, v := range mgr.tasks {
			if v.TaskID == taskid {
				mgr.tasks = append(mgr.tasks[:i], mgr.tasks[i+1:]...)

				break
			}
		}
	}

	if len(mgr.tasks) == 0 {
		mainLogger.Debug("nextTask: no task")

		return true
	}

	failnums := 0
	for _, v := range mgr.tasks {
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

	if failnums == len(mgr.tasks) {
		mainLogger.Debug("nextTask: task fail", zap.Int("tailnums", failnums))

		return true
	}

	return false
}

// newTaskID -
func (mgr *ClientMgr) newTaskID() int {
	mgr.MaxTaskID++

	return mgr.MaxTaskID
}

// GetConfig - get config
func (mgr *ClientMgr) GetConfig() Config {
	return *mgr.cfg
}

// GetAnkaDB - get ankaDB
func (mgr *ClientMgr) GetAnkaDB() ankadb.AnkaDB {
	return mgr.db.ankaDB
}
