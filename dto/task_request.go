package dto

import (
	"time"

	"gorm.io/datatypes"
)

type CreateTaskRequest struct {
	Name     string         `json:"name" binding:"required"`
	DueDate  *time.Time     `json:"due_date,omitempty"`
	Assignee string         `json:"assignee,omitempty"`
	Tags     datatypes.JSON `json:"tags,omitempty"`
}

type UpdateTaskRequest struct {
	Name     *string         `json:"name,omitempty"`
	Status   *int            `json:"status,omitempty"` // 0 = 未完成，1 = 已完成
	DueDate  *time.Time      `json:"due_date,omitempty"`
	Assignee *string         `json:"assignee,omitempty"`
	Tags     *datatypes.JSON `json:"tags,omitempty"`
}
