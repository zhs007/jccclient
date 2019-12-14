package jccclient

// PipelineCode - pipeline code
type PipelineCode string

// Pipeline - pipeline
type Pipeline struct {
	Tasks         []*Task
	Running       []*Task
	Code          PipelineCode
	IsMulHostname bool
}

// AddTask - add a task
func (p *Pipeline) AddTask(t *Task) {
	p.Tasks = append(p.Tasks, t)
}

// PipelineMgr - prpilines manager
type PipelineMgr struct {
	Pipelines map[PipelineCode]*Pipeline
}

// newPipelineMgr - new PipelineMgr
func newPipelineMgr() *PipelineMgr {
	return &PipelineMgr{
		Pipelines: make(map[PipelineCode]*Pipeline),
	}
}
