package handler

import (
	"net/http"
	"task-api/model"

	"task-api/repository"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	repo repository.TaskRepository
}

func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a task with name and status
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   task body model.Task true "Task to create"
// @Success 201 {object} model.Task
// @Failure 400 {object} gin.H
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}
