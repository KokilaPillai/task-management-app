package handlers

import (
	"log"

	"github.com/ranefattesingh/task-management-app/server/repo"
)

type task struct {
	l *log.Logger
	r repo.Repository
}

func NewTaskHandler(l *log.Logger, r repo.Repository) *task {
	return &task{l, r}
}
