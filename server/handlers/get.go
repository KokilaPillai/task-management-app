package handlers

import (
	"net/http"
	"tasktracker/server/data"
)

func (h *task) GetAll(rw http.ResponseWriter, r *http.Request) {
	tl, err := data.GetTasks()
	if err != nil {
		http.Error(rw, data.JsonError(data.ErrFailedToGet), http.StatusInternalServerError)
		return
	}

	if err := tl.ToJSON(rw); err != nil {
		http.Error(rw, data.JsonError(data.ErrStructToJson), http.StatusInternalServerError)
		return
	}
}
