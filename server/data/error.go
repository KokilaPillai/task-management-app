package data

import (
	"encoding/json"
	"errors"
)

type Error struct {
	Message string `json:"message"`
}

func JsonError(s error) string {
	e := &Error{
		Message: s.Error(),
	}
	d, _ := json.Marshal(e)
	return string(d)
}

var ErrJsonToStruct = errors.New("Error while converting request Json to type Task")
var ErrStructToJson = errors.New("Error while converting response to Json")
var ErrFailedToInsert = errors.New("Error while inserting task")
var ErrFailedToGet = errors.New("Error while getting list of tasks")
var ErrFailedToUpdate = errors.New("Error while updating task")
var ErrInvalidUrlParameter = errors.New("Error invalid id received")
var ErrTaskNotFound = errors.New("Task not found")
