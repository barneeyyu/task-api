package dto

import (
	"time"
)

type TaskResponse struct {
	ID        uint       `json:"id" example:"1"`
	Name      string     `json:"name" example:"write a blog"`
	Status    int        `json:"status" example:"1"`
	DueDate   *time.Time `json:"due_date,omitempty" example:"2025-06-20T10:00:00Z"`
	Assignee  string     `json:"assignee" example:"Barney"`
	Tags      []string   `json:"tags,omitempty" example:"[\"doc\",\"internal\",\"urgent\"]"`
	CreatedAt time.Time  `json:"created_at" example:"2025-06-20T10:00:00Z"`
	UpdatedAt time.Time  `json:"updated_at" example:"2025-06-20T10:00:00Z"`
}
