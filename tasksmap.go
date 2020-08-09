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

// DelTask - delete a task
func (m *TasksMap) DelTask(key string, tid int) {
	lst, isok := m.MapTasks[key]
	if !isok {
		return
	}

	if len(lst) == 1 && lst[0].TaskID == tid {
		delete(m.MapTasks, key)

		return
	}

	for i, v := range lst {
		if v.TaskID == tid {
			lst = append(lst[:i], lst[i+1:]...)

			break
		}
	}

	m.MapTasks[key] = lst
}

// DelTaskEx - delete a task
func (m *TasksMap) DelTaskEx(tid int) {
	for k, lst := range m.MapTasks {
		if len(lst) == 1 && lst[0].TaskID == tid {
			delete(m.MapTasks, k)

			return
		}

		for i, v := range lst {
			if v.TaskID == tid {
				lst = append(lst[:i], lst[i+1:]...)

				break
			}
		}

		m.MapTasks[k] = lst
	}
}
