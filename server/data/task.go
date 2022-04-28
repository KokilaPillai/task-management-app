package data

import (
	"encoding/json"
	"io"
)

type Task struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
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

func AddTask(t *Task) (*Task, error) {
	t.ID = taskList.getNextID()
	taskList = append(taskList, t)
	return taskList[len(taskList)-1], nil
}

func GetTasks() (Tasks, error) {
	return taskList, nil
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
