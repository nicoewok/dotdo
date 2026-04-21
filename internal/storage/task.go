package storage

import "time"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"` // "todo", "doing", "done"
	Created   time.Time `json:"created"`
	Completed time.Time `json:"completed,omitempty"`
}

type List struct {
	Tasks []Task `json:"tasks"`
}
