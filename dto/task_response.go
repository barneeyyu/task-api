package dto

import (
	"time"

	"gorm.io/datatypes"
)

type TaskResponse struct {
	ID        uint            `json:"id"`
	Name      string          `json:"name"`
	Status    int             `json:"status"`
	DueDate   *time.Time      `json:"due_date,omitempty"`
	Assignee  string          `json:"assignee"`
	Tags      *datatypes.JSON `json:"tags,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
