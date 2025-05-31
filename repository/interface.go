package repository

import "task-api/model"

type RepositoryInterface interface {
	CreateTask(task *model.Task) (*model.Task, error)
	GetTaskByID(id uint) (*model.Task, error)
	GetAllTasks() ([]model.Task, error)
	UpdateTask(fields map[string]interface{}, id uint) error
	DeleteTask(id uint) (bool, error)
}
