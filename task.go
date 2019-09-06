package jccclient

import (
	"context"

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

// Task - task
type Task struct {
	Callback      FuncTaskCallback
	AnalyzePage   *TaskAnalyzePage
	GeoIP         *TaskGeoIP
	TechInAsia    *TaskTechInAsia
	SteepAndCheap *TaskSteepAndCheap
	Timeout       int
	RetryNums     int
	tags          *Tags
	hostname      string
	taskid        int
	running       bool
	fail          bool
}
