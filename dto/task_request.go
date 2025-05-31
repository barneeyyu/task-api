package dto

import (
	"time"
)

type CreateTaskRequest struct {
	Name     string     `json:"name" binding:"required,max=100"  example:"write a blog"`
	DueDate  *time.Time `json:"due_date,omitempty"                 example:"2025-06-20T10:00:00Z"`
	Assignee string     `json:"assignee,omitempty" binding:"max=10" example:"Barney"`
	Tags     []string   `json:"tags,omitempty" binding:"max=3,dive,max=10" example:"[\"doc\",\"internal\",\"urgent\"]"`
}

type UpdateTaskRequest struct {
	Name     *string    `json:"name,omitempty"     binding:"omitempty,max=100"   example:"write a blog"`
	Status   *int       `json:"status,omitempty"   binding:"omitempty,oneof=0 1" example:"1"` // 0 或 1
	DueDate  *time.Time `json:"due_date,omitempty"                                      example:"2025-06-20T10:00:00Z"`
	Assignee *string    `json:"assignee,omitempty" binding:"omitempty,max=10"     example:"Barney"`
	Tags     *[]string  `json:"tags,omitempty"     binding:"omitempty,max=3,dive,max=10" example:"[\"doc\",\"internal\",\"urgent\"]"`
}
