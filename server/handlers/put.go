package handlers

import (
	"net/http"
	"strconv"

	"github.com/ranefattesingh/task-management-app/server/data"

	"github.com/gorilla/mux"
)

func (h *task) UpdateTask(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("[INFO]\tReceived UpdateTasks")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, data.JsonError(data.ErrInvalidUrlParameter), http.StatusInternalServerError)
		return
	}

	req := r.Context().Value(&Request{}).(*data.Task)

	res, err := h.r.UpdateTask(id, req)
	if err != nil {
		if err == data.ErrTaskNotFound {
			http.Error(rw, data.JsonError(err), http.StatusNotFound)
			return
		}

		http.Error(rw, data.JsonError(data.ErrFailedToUpdate), http.StatusInternalServerError)
		return
	}

	if err := res.ToJSON(rw); err != nil {
		http.Error(rw, data.JsonError(data.ErrStructToJson), http.StatusInternalServerError)
		return
	}
}
