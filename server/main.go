package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ranefattesingh/task-management-app/server/handlers"
	"github.com/ranefattesingh/task-management-app/server/repo"
	"github.com/ranefattesingh/task-management-app/server/repo/mock"
	"github.com/ranefattesingh/task-management-app/server/repo/sql"

	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	PORT              = 5000
	FRONTEND          = 4200
	MOCK              = false
	CONNECTION_POOL   = 10
	CONNECTION_STRING = "./Task.db"
)

// "Origin", "Accept", "Content-Type", "X-Requested-With"
var AllowedHeaders = []string{"Content-Type", "X-Requested-With"}
var AllowedOrigins = []string{fmt.Sprintf("http://localhost:%v", FRONTEND)}
var AllowedMethods = []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}

func main() {

	l := log.New(os.Stdout, "TaskAPI: ", log.LstdFlags)

	var r repo.Repository
	if MOCK {
		r = mock.NewMock()
	} else {
		r = sql.NewSql(CONNECTION_STRING, CONNECTION_POOL)
	}

	m := mux.NewRouter()
	th := handlers.NewTaskHandler(l, r)

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

	// CORS
	headersOk := gohandlers.AllowedHeaders(AllowedHeaders)
	originsOk := gohandlers.AllowedOrigins(AllowedOrigins)
	methodsOk := gohandlers.AllowedMethods(AllowedMethods)
	ch := gohandlers.CORS(originsOk, headersOk, methodsOk)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%v", PORT),
		Handler:      ch(m),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	signal := make(chan os.Signal)

	l.Println("Received terminate, graceful shutdown", <-signal)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
