package jccclient

// RebuildTasks - resort tasks with hostname
func RebuildTasks(tasks []*Task) []*Task {
	m := NewTasksMap()

	for _, v := range tasks {
		m.AddTask(v.Hostname, v)
	}

	nl := []*Task{}

	return nl
}
