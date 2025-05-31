package repository

import (
	"errors"
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

func (r *TaskRepository) UpdateTask(fields map[string]interface{}, id uint) error {
	result := r.db.Model(&model.Task{}).
		Where("id = ?", id).
		Updates(fields)

	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (r *TaskRepository) DeleteTask(id uint) (bool, error) {
	result := r.db.Delete(&model.Task{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
