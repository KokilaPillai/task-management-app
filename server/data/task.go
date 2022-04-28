package data

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

type Task struct {
	ID       int    `json:"id"`
	Text     string `json:"text" validate:"required,gte=1"`
	Day      string `json:"day"`
	Reminder bool   `json:"reminder"`
}

func (t *Task) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(t)
}

func (t *Task) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(t)
}

type Tasks []*Task

func (tl *Tasks) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(tl)
}

func (tl Tasks) getNextID() int {
	if len(tl) == 0 {
		return 1
	}

	t := tl[len(tl)-1]
	return t.ID + 1
}

func (t *Task) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

func AddTask(t *Task) (*Task, error) {
	t.ID = taskList.getNextID()
	taskList = append(taskList, t)
	return taskList[len(taskList)-1], nil
}

func GetTasks() (Tasks, error) {
	return taskList, nil
}

func findIndexByTaskId(id int) int {
	for i, item := range taskList {
		if item.ID == id {
			return i
		}
	}
	return -1
}

func UpdateTask(id int, t *Task) (*Task, error) {
	index := findIndexByTaskId(id)

	if index == -1 {
		return nil, ErrTaskNotFound
	}

	taskList[index].Text = t.Text
	taskList[index].Day = t.Day
	taskList[index].Reminder = t.Reminder

	return taskList[index], nil

}

func DeleteTask(id int) error {
	index := findIndexByTaskId(id)

	if index == -1 {
		return ErrTaskNotFound
	}

	taskList = append(taskList[:index], taskList[index+1:]...)
	return nil
}

func GetTask(id int) (*Task, error) {
	index := findIndexByTaskId(id)

	if index == -1 {
		return nil, ErrTaskNotFound
	}

	return taskList[index], nil
}

var taskList = Tasks{
	{
		ID:       1,
		Text:     "Buy PanCake",
		Day:      "27th May 2022",
		Reminder: true,
	},
	{
		ID:       2,
		Text:     "Buy Bananas",
		Day:      "25th May 2022",
		Reminder: false,
	},
}
