package jccclient

import (
	"context"
	"encoding/json"

	jarviscrawlercore "github.com/zhs007/jccclient/proto"
)

// FuncTaskCallback - callback
type FuncTaskCallback func(context.Context, *Task, error, *jarviscrawlercore.ReplyCrawler)

// TaskAnalyzePage - task analyzepage
type TaskAnalyzePage struct {
	URL      string
	Viewport Viewport
	Options  AnalyzePageOptions
}

// TaskGeoIP - task get ip geolocation
type TaskGeoIP struct {
	IP       string
	Platform string
}

// TaskTechInAsia - task get data in techinasia
type TaskTechInAsia struct {
	Mode        jarviscrawlercore.TechInAsiaMode
	JobCode     string
	CompanyCode string
	JobNums     int
	JobTag      string
	JobSubTag   string
}

// TaskSteepAndCheap - task get data in steepandcheap
type TaskSteepAndCheap struct {
	Mode jarviscrawlercore.SteepAndCheapMode
	URL  string
	Page int
}

// TaskJRJ - task get data in jrj
type TaskJRJ struct {
	Mode jarviscrawlercore.JRJMode
	Code string
	Year string
}

// TaskJD - task get data in jd
type TaskJD struct {
	Mode jarviscrawlercore.JDMode
	URL  string
}

// TaskAlimama - task get data in alimama
type TaskAlimama struct {
	Mode jarviscrawlercore.AlimamaMode
	Text string
	URL  string
}

// TaskTmall - task get data in tmall
type TaskTmall struct {
	Mode   jarviscrawlercore.TmallMode
	ItemID string
	Device string
}

// TaskTaobao - task get data in taobao
type TaskTaobao struct {
	Mode   jarviscrawlercore.TaobaoMode
	ItemID string
	Device string
	Text   string
}

// TaskMountainSteals - task get data in mountainsteals
type TaskMountainSteals struct {
	Mode jarviscrawlercore.MountainStealsMode
	URL  string
}

// TaskDouban - task get data in douban
type TaskDouban struct {
	Mode       jarviscrawlercore.DoubanMode
	Text       string
	ID         string
	DoubanType jarviscrawlercore.DoubanType
}

// TaskManhuaDB - task get data in manhuadb
type TaskManhuaDB struct {
	Mode     jarviscrawlercore.ManhuaDBMode
	AuthorID string
}

// TaskOABT - task get data in oabt
type TaskOABT struct {
	Mode      jarviscrawlercore.OABTMode
	PageIndex int32
}

// TaskHao6v - task get data in hao6v
type TaskHao6v struct {
	Mode jarviscrawlercore.Hao6VMode
	URL  string
}

// Task - task
type Task struct {
	Callback       FuncTaskCallback    `json:"-"`
	AnalyzePage    *TaskAnalyzePage    `json:"AnalyzePage"`
	GeoIP          *TaskGeoIP          `json:"GeoIP"`
	TechInAsia     *TaskTechInAsia     `json:"TechInAsia"`
	SteepAndCheap  *TaskSteepAndCheap  `json:"SteepAndCheap"`
	JRJ            *TaskJRJ            `json:"JRJ"`
	JD             *TaskJD             `json:"JD"`
	Alimama        *TaskAlimama        `json:"Alimama"`
	Tmall          *TaskTmall          `json:"Tmall"`
	Taobao         *TaskTaobao         `json:"Taobao"`
	MountainSteals *TaskMountainSteals `json:"mountainsteals"`
	Douban         *TaskDouban         `json:"douban"`
	ManhuaDB       *TaskManhuaDB       `json:"manhuadb"`
	OABT           *TaskOABT           `json:"oabt"`
	Hao6v          *TaskHao6v          `json:"hao6v"`
	Timeout        int                 `json:"timeout"`
	RetryNums      int                 `json:"retrynums"`
	Tags           *Tags               `json:"tags"`
	Hostname       string              `json:"hostname"`
	TaskID         int                 `json:"taskID"`
	Running        bool                `json:"running"`
	Fail           bool                `json:"fail"`
	ServAddr       string              `json:"servaddr"`
}

// Reset - reset
func (task *Task) Reset() {
	task.Fail = false
	task.Running = false
}

// ToString - to string
func (task *Task) ToString() string {
	result, err := json.Marshal(task)
	if err != nil {
		return err.Error()
	}

	return string(result)
}

// GetPipelineCode - get pipeline code
func (task *Task) GetPipelineCode() (PipelineCode, bool) {
	if task.Tmall != nil || task.Taobao != nil {
		return "taobao.com", true
	}

	return PipelineCode(task.Hostname), false
}
