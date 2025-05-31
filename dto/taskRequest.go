package dto

import (
	"time"

	"gorm.io/datatypes"
)

type TaskRequest struct {
	Name     string         `json:"name" binding:"required"`
	DueDate  *time.Time     `json:"due_date,omitempty"`
	Assignee string         `json:"assignee,omitempty"`
	Tags     datatypes.JSON `json:"tags,omitempty"`
}
