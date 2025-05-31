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

type TaskHandler struct {
	repo repository.RepositoryInterface
}

func NewTaskHandler(repo repository.RepositoryInterface) *TaskHandler {
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
// @Failure      400 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var request dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	task := model.Task{
		Name:     request.Name,
		Status:   0, // 預設未完成
		DueDate:  request.DueDate,
		Assignee: request.Assignee,
		Tags:     request.Tags,
	}

	createdTask, err := h.repo.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
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
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /tasks [get]
// @Router       /tasks/{id} [get]
func (h *TaskHandler) GetTasks(c *gin.Context) {
	idStr := c.Query("id")
	if idStr != "" {
		idUint, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid id format"})
			return
		}
		task, err := h.repo.GetTaskByID(uint(idUint))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "task not found"})
			} else {
				c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, dto.TaskResponse(*task))
		return
	}

	tasks, err := h.repo.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
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
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid id format"})
		return
	}

	var request dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	fields := map[string]interface{}{}
	if request.Name != nil {
		fields["name"] = *request.Name
	}
	if request.Status != nil {
		fields["status"] = *request.Status
	}
	if request.DueDate != nil {
		fields["due_date"] = *request.DueDate
	}
	if request.Assignee != nil {
		fields["assignee"] = *request.Assignee
	}
	if request.Tags != nil {
		fields["tags"] = request.Tags
	}

	if err := h.repo.UpdateTask(fields, uint(idUint)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteTask godoc
// @Summary      Delete a task
// @Description  Delete a task by ID
// @Tags         tasks
// @Produce      json
// @Param        id path int true "Task ID"
// @Success      204 "No Content"
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid id format"})
		return
	}

	deleted, err := h.repo.DeleteTask(uint(idUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "task not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
