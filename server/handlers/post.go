package handlers

import (
	"net/http"

	"github.com/ranefattesingh/task-management-app/server/data"
)

func (h *task) AddTask(rw http.ResponseWriter, r *http.Request) {

	req := r.Context().Value(&Request{}).(*data.Task)

	res, err := data.AddTask(req)
	if err != nil {
		http.Error(rw, data.JsonError(data.ErrFailedToInsert), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	if err := res.ToJSON(rw); err != nil {
		http.Error(rw, data.JsonError(data.ErrStructToJson), http.StatusInternalServerError)
		return
	}
}
