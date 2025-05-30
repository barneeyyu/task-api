package test

import (
	"testing"
	"time"

	"task-api/model"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	task := model.Task{
		Name:   "Test Task",
		Status: 0,
	}
	assert.Equal(t, "Test Task", task.Name)
	assert.Equal(t, 0, task.Status)
	assert.WithinDuration(t, time.Now(), task.CreatedAt, time.Second) // 預期時間接近
}
