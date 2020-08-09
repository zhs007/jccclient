package jccclient

// RebuildTasks - resort tasks with hostname
func RebuildTasks(tasks []*Task) []*Task {
	m := NewTasksMap()

	for _, v := range tasks {
		m.AddTask(v.Hostname, v)
	}

	nl := []*Task{}

	for true {
		for k, v := range m.MapTasks {
			t := v[0]

			nl = append(nl, t)

			if len(v) > 1 {
				m.MapTasks[k] = v[1:]
			} else {
				delete(m.MapTasks, k)
			}
		}

		if len(m.MapTasks) == 0 {
			break
		}
	}

	return nl
}
