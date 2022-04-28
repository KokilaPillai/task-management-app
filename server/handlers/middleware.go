package handlers

import (
	"context"
	"net/http"
	"tasktracker/server/data"
)

type Request struct{}

func (t *task) RequestValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		req := &data.Task{}
		if err := req.FromJSON(r.Body); err != nil {
			http.Error(rw, data.JsonError(data.ErrJsonToStruct), http.StatusBadRequest)
			return
		}

		// Perform Validation here

		r = r.WithContext(context.WithValue(r.Context(), &Request{}, req))

		next.ServeHTTP(rw, r)
	})
}