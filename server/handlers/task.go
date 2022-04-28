package handlers

import "log"

type task struct {
	l *log.Logger
}

func NewTaskHandler(l *log.Logger) *task {
	return &task{l}
}
