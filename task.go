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
}

// Task - task
type Task struct {
	Callback    FuncTaskCallback
	AnalyzePage *TaskAnalyzePage
	GeoIP       *TaskGeoIP
	TechInAsia  *TaskTechInAsia
	Timeout     int
	RetryNums   int
	tags        *Tags
	hostname    string
	taskid      int
	running     bool
}
