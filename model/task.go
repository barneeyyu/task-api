package model

import "time"

type Task struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // 任務名稱
	Status    int       `gorm:"type:integer;not null" json:"status"`    // 0 = 未完成，1 = 已完成
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
