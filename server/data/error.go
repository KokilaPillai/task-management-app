package data

import "encoding/json"

type Error struct {
	Message string `json:"message"`
}

func JsonError(s string) string {
	e := &Error{
		Message: s,
	}
	d, _ := json.Marshal(e)
	return string(d)
}

var ErrJsonToStruct = "Error while converting request Json to type Task"
var ErrStructToJson = "Error while converting response to Json"
var ErrFailedToInsert = "Error while inserting task"
var ErrFailedToGet = "Error while getting list of tasks"
var ErrFailedToUpdate = "Error while updating task"
var ErrInvalidUrlParameter = "Error invalid id received"
