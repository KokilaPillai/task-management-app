package mock

import (
	"log"

	"github.com/ranefattesingh/task-management-app/server/data"
)

type mock struct {
	l *log.Logger
}

func NewMock() *mock {
	return &mock{}
}

func getNextID() uint {
	if len(taskList) == 0 {
		return 1
	}

	t := taskList[len(taskList)-1]
	return t.ID + 1
}

func (*mock) AddTask(t *data.Task) (*data.Task, error) {
	t.ID = getNextID()
	taskList = append(taskList, t)
	return taskList[len(taskList)-1], nil
}

func (*mock) GetTasks() (data.Tasks, error) {
	return taskList, nil
}

func findIndexByTaskId(id uint) int {
	for i, item := range taskList {
		if item.ID == id {
			return i
		}
	}
	return -1
}

func (*mock) UpdateTask(id int, t *data.Task) (*data.Task, error) {
	index := findIndexByTaskId(uint(id))

	if index == -1 {
		return nil, data.ErrTaskNotFound
	}

	taskList[index].Text = t.Text
	taskList[index].Day = t.Day
	taskList[index].Reminder = t.Reminder

	return taskList[index], nil

}

func (*mock) DeleteTask(id int) error {
	index := findIndexByTaskId(uint(id))

	if index == -1 {
		return data.ErrTaskNotFound
	}

	taskList = append(taskList[:index], taskList[index+1:]...)
	return nil
}

func (*mock) GetTask(id int) (*data.Task, error) {
	index := findIndexByTaskId(uint(id))

	if index == -1 {
		return nil, data.ErrTaskNotFound
	}

	return taskList[index], nil
}

var taskList = data.Tasks{}
