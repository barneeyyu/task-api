package repository

import "task-api/model"

type RepositoryInterface interface {
	CreateTask(task model.Task) error
}
