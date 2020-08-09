package jccclient

// TasksMap - tasks map
type TasksMap struct {
	MapTasks map[string]([]*Task)
}

// NewTasksMap - new TasksMap
func NewTasksMap() *TasksMap {
	return &TasksMap{
		MapTasks: make(map[string]([]*Task)),
	}
}

// AddTask - add a task
func (m *TasksMap) AddTask(key string, t *Task) {
	_, isok := m.MapTasks[key]
	if isok {
		m.MapTasks[key] = append(m.MapTasks[key], t)
	} else {
		m.MapTasks[key] = []*Task{t}
	}
}

// Count - count
func (m *TasksMap) Count() int {
	nums := 0

	for _, v := range m.MapTasks {
		nums += len(v)
	}

	return nums
}
