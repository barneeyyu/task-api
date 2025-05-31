package repository

import (
	"task-api/model"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *model.Task) (*model.Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) GetTaskByID(id uint) (*model.Task, error) {
	var task model.Task
	if err := r.db.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
