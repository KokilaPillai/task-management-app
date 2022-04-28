package handlers

import (
	"net/http"
	"strconv"

	"github.com/ranefattesingh/task-management-app/server/data"

	"github.com/gorilla/mux"
)

func (h *task) DeleteTask(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("[INFO]\tReceived DeleteTask")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, data.JsonError(data.ErrInvalidUrlParameter), http.StatusInternalServerError)
		return
	}

	err = h.r.DeleteTask(id)
	if err != nil {
		if err == data.ErrTaskNotFound {
			http.Error(rw, data.JsonError(err), http.StatusNotFound)
			return
		}

		http.Error(rw, data.JsonError(data.ErrFailedToUpdate), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
