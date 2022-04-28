package handlers

import (
	"net/http"
	"strconv"
	"tasktracker/server/data"

	"github.com/gorilla/mux"
)

func (h *task) GetTasks(rw http.ResponseWriter, r *http.Request) {
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

func (h *task) GetTask(rw http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, data.JsonError(data.ErrInvalidUrlParameter), http.StatusInternalServerError)
		return
	}

	t, err := data.GetTask(id)
	if err != nil {
		if err == data.ErrTaskNotFound {
			http.Error(rw, data.JsonError(err), http.StatusNotFound)
			return
		}

		http.Error(rw, data.JsonError(data.ErrTaskNotFound), http.StatusInternalServerError)
		return
	}

	if err := t.ToJSON(rw); err != nil {
		http.Error(rw, data.JsonError(data.ErrStructToJson), http.StatusInternalServerError)
		return
	}
}
