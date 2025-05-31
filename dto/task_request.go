package dto

import (
	"time"
)

type CreateTaskRequest struct {
	Name     string     `json:"name" binding:"required" example:"write a blog"`
	DueDate  *time.Time `json:"due_date,omitempty" example:"2025-06-20T10:00:00Z"`
	Assignee string     `json:"assignee,omitempty" example:"Barney"`
	Tags     []string   `json:"tags,omitempty" example:"[\"doc\",\"internal\",\"urgent\"]"`
}

type UpdateTaskRequest struct {
	Name     *string    `json:"name,omitempty" example:"write a blog"`
	Status   *int       `json:"status,omitempty" example:"1"` // 0 = 未完成，1 = 已完成
	DueDate  *time.Time `json:"due_date,omitempty" example:"2025-06-20T10:00:00Z"`
	Assignee *string    `json:"assignee,omitempty" example:"Barney"`
	Tags     *[]string  `json:"tags,omitempty" example:"[\"doc\",\"internal\",\"urgent\"]"`
}
