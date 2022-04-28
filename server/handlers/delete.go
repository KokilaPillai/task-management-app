package handlers

import (
	"net/http"
	"strconv"
	"tasktracker/server/data"

	"github.com/gorilla/mux"
)

func (h *task) DeleteTask(rw http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, data.JsonError(data.ErrInvalidUrlParameter), http.StatusInternalServerError)
		return
	}

	err = data.DeleteTask(id)
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
