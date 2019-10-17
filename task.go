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

// Task - task
type Task struct {
	Callback      FuncTaskCallback   `json:"-"`
	AnalyzePage   *TaskAnalyzePage   `json:"AnalyzePage"`
	GeoIP         *TaskGeoIP         `json:"GeoIP"`
	TechInAsia    *TaskTechInAsia    `json:"TechInAsia"`
	SteepAndCheap *TaskSteepAndCheap `json:"SteepAndCheap"`
	JRJ           *TaskJRJ           `json:"JRJ"`
	Timeout       int                `json:"timeout"`
	RetryNums     int                `json:"retrynums"`
	Tags          *Tags              `json:"tags"`
	Hostname      string             `json:"hostname"`
	TaskID        int                `json:"taskID"`
	Running       bool               `json:"running"`
	Fail          bool               `json:"fail"`
	ServAddr      string             `json:"servaddr"`
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
