package Struct

import "time"

type Task struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Status    string     `json:"status"` // "todo", "in-progress", "done"
	CreateAT  time.Time  `json:"createdAT"`
	UpdateAT  time.Time  `json:"updateAT"`
	CompleteAT *time.Time `json:"completeAT"`
}
