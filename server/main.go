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
	get.HandleFunc("/tasks", th.GetTasks)
	get.HandleFunc("/tasks/{id:[0-9]+}", th.GetTask)

	post := m.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/tasks", th.AddTask)
	post.Use(th.RequestValidationMiddleware)

	put := m.Methods(http.MethodPut).Subrouter()
	put.HandleFunc("/tasks/{id:[0-9]+}", th.UpdateTask)
	put.Use(th.RequestValidationMiddleware)

	delete := m.Methods(http.MethodDelete).Subrouter()
	delete.HandleFunc("/tasks/{id:[0-9]+}", th.DeleteTask)

	s := &http.Server{
		Addr:    ":4001",
		Handler: m,
	}

	s.ListenAndServe()
	// http.ListenAndServe(":4001", nil)
}
