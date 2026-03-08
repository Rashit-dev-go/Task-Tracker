package main

import (
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Status      Status     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type TaskStorage struct {
	Tasks []Task `json:"tasks"`
}
