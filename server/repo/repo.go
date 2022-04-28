package repo

import "github.com/ranefattesingh/task-management-app/server/data"

type Repository interface {
	AddTask(t *data.Task) (*data.Task, error)
	GetTasks() (data.Tasks, error)
	UpdateTask(id int, t *data.Task) (*data.Task, error)
	DeleteTask(id int) error
	GetTask(id int) (*data.Task, error)
}
