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

var ErrJsonToStruct = "Error while converting request Json to type Task."
var ErrStructToJson = "Error while converting response to Json."
var ErrFailedToInsert = "Error while trying to insert task in database."
var ErrFailedToGet = "Error while trying to read data from database."
