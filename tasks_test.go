package jccclient

import "testing"

func Test_RebuildTasks(t *testing.T) {
	lst := []*Task{
		&Task{Hostname: "a", TaskID: 1},
		&Task{Hostname: "a", TaskID: 2},
		&Task{Hostname: "a", TaskID: 3},
		&Task{Hostname: "b", TaskID: 11},
		&Task{Hostname: "b", TaskID: 12},
		&Task{Hostname: "c", TaskID: 21},
		&Task{Hostname: "a", TaskID: 4},
		&Task{Hostname: "c", TaskID: 22},
		&Task{Hostname: "c", TaskID: 23},
		&Task{Hostname: "c", TaskID: 24},
	}

	lst1 := [][]*Task{
		[]*Task{
			&Task{Hostname: "a", TaskID: 1},
			&Task{Hostname: "b", TaskID: 11},
			&Task{Hostname: "c", TaskID: 21},
		},
		[]*Task{
			&Task{Hostname: "a", TaskID: 2},
			&Task{Hostname: "b", TaskID: 12},
			&Task{Hostname: "c", TaskID: 22},
		},
		[]*Task{
			&Task{Hostname: "a", TaskID: 3},
			&Task{Hostname: "c", TaskID: 23},
		},
		[]*Task{
			&Task{Hostname: "a", TaskID: 4},
			&Task{Hostname: "c", TaskID: 24},
		},
	}

	lst2 := RebuildTasks(lst)

	for j, v := range lst2 {
		isok := false

		lst3 := []*Task{}

		for _, v1 := range lst1[0] {
			if !(v.Hostname == v1.Hostname && v.TaskID == v1.TaskID) {
				lst3 = append(lst3, v1)
			} else {
				isok = true
			}
		}

		if !isok {
			t.Fatalf("RebuildTasks %v err [%s,%d]", j,
				v.Hostname, v.TaskID)
		}

		if len(lst3) == 0 {
			lst1 = lst1[1:]
		} else {
			lst1[0] = lst3
		}
	}

	t.Log("Test_RebuildTasks OK")
}
