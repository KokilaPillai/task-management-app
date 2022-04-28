package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"tasktracker/server/handlers"
	"time"

	"github.com/gorilla/mux"
)

const PORT = 5000

func main() {

	l := log.New(os.Stdout, "TaskAPI: ", log.LstdFlags)

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
		Addr:         fmt.Sprintf(":%v", PORT),
		Handler:      m,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		s.ListenAndServe()
	}()

	signal := make(chan os.Signal)

	l.Println("Received terminate, graceful shutdown", <-signal)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
