package handlers

import (
	"net/http"
	"strconv"
	"tasktracker/server/data"

	"github.com/gorilla/mux"
)

func (h *task) UpdateTask(rw http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, data.JsonError(data.ErrInvalidUrlParameter), http.StatusInternalServerError)
		return
	}

	req := r.Context().Value(&Request{}).(*data.Task)

	res, err := data.UpdateTask(id, req)
	if err != nil {
		http.Error(rw, data.JsonError(data.ErrFailedToUpdate), http.StatusInternalServerError)
		return
	}

	if err := res.ToJSON(rw); err != nil {
		http.Error(rw, data.JsonError(data.ErrStructToJson), http.StatusInternalServerError)
		return
	}
}
