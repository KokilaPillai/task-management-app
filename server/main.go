package main

import (
	"net/http"
	"tasktracker/server/handlers"

	"github.com/gorilla/mux"
)

func main() {

	m := mux.NewRouter()
	th := handlers.NewTaskHandler()

	get := m.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("/tasks", th.GetAll)

	post := m.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/tasks", th.AddTask)

	put := m.Methods(http.MethodPut).Subrouter()
	put.HandleFunc("/tasks", th.UpdateTask)

	s := &http.Server{
		Addr:    ":4001",
		Handler: m,
	}

	s.ListenAndServe()
	// http.ListenAndServe(":4001", nil)
}
