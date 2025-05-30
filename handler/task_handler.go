package handler

import (
	"net/http"
	"task-api/model"

	"task-api/repository"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a generic error message
type ErrorResponse struct {
	Error string `json:"error"`
}

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
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := h.repo.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
	}

	c.JSON(http.StatusCreated, task)
}
