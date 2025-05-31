package model

import (
	"time"
)

type Task struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"size:255;not null" json:"name"`
	Status    int        `gorm:"type:int;default:0" json:"status"` // 0 = 未完成，1 = 已完成
	DueDate   *time.Time `json:"due_date,omitempty"`
	Assignee  string     `json:"assignee"`
	Tags      []string   `gorm:"type:json;serializer:json" json:"tags,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
