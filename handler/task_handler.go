package handler

import (
	"errors"
	"net/http"
	"strconv"

	"task-api/dto"
	"task-api/model"
	"task-api/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
// @Summary      Create a new task
// @Description  Create a task with name, due date, assignee and tags
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task body dto.CreateTaskRequest true "Task to create"
// @Success      201 {object} dto.TaskResponse
// @Failure      400 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var request dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	task := model.Task{
		Name:     request.Name,
		Status:   0, // 預設未完成
		DueDate:  request.DueDate,
		Assignee: request.Assignee,
		Tags:     &request.Tags,
	}

	createdTask, err := h.repo.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	response := dto.TaskResponse(*createdTask)
	c.JSON(http.StatusCreated, response)
}

// GetTasks godoc
// @Summary      Get task(s)
// @Description  Get a single task by ID (if provided), or all tasks
// @Tags         tasks
// @Produce      json
// @Param        id query int false "Task ID"
// @Success      200 {object} dto.TaskResponse
// @Success      200 {array} dto.TaskResponse
// @Failure      400 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /tasks [get]
func (h *TaskHandler) GetTasks(c *gin.Context) {
	idStr := c.Query("id")
	if idStr != "" {
		idUint, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid id format"})
			return
		}
		task, err := h.repo.GetTaskByID(uint(idUint))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, ErrorResponse{Error: "task not found"})
			} else {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, dto.TaskResponse(*task))
		return
	}

	tasks, err := h.repo.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	var responses []dto.TaskResponse
	for _, task := range tasks {
		responses = append(responses, dto.TaskResponse(task))
	}

	c.JSON(http.StatusOK, responses)
}

// UpdateTask godoc
// @Summary      Update a task
// @Description  Update task fields by ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id path int true "Task ID"
// @Param        task body dto.UpdateTaskRequest true "Updated task data"
// @Success      200 {object} dto.TaskResponse
// @Failure      400 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid id format"})
		return
	}

	var request dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	task, err := h.repo.GetTaskByID(uint(idUint))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: "task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}
		return
	}

	// 更新欄位
	if request.Name != nil {
		task.Name = *request.Name
	}
	if request.Status != nil {
		task.Status = *request.Status
	}
	if request.DueDate != nil {
		task.DueDate = request.DueDate
	}
	if request.Assignee != nil {
		task.Assignee = *request.Assignee
	}
	if request.Tags != nil {
		task.Tags = request.Tags
	}

	if err := h.repo.UpdateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.TaskResponse(*task))
}
