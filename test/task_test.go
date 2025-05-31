package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"task-api/dto"
	"task-api/handler"
	"task-api/model"
	"task-api/repository"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// mockRepo implements TaskRepository for testing
type mockRepo struct{}

func (m *mockRepo) CreateTask(task *model.Task) (*model.Task, error) {
	task.ID = 1
	return task, nil
}

func (m *mockRepo) GetTaskByID(id uint) (*model.Task, error) {
	if id == 1 {
		return &testTask, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockRepo) GetAllTasks() ([]model.Task, error) {
	return []model.Task{testTask}, nil
}

func (m *mockRepo) UpdateTask(fields map[string]interface{}, id uint) error {
	if id != 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *mockRepo) DeleteTask(id uint) (bool, error) {
	if id == 1 {
		return true, nil
	}
	return false, nil
}

// 確保 mockRepo 符合 interface，放在 mockRepo 定義後
var _ repository.RepositoryInterface = (*mockRepo)(nil)

var testTask = model.Task{
	ID:        1,
	Name:      "Test Task",
	Status:    0,
	Assignee:  "Tester",
	Tags:      []string{"test"},
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := handler.NewTaskHandler(&mockRepo{})

	r.POST("/tasks", h.CreateTask)
	r.GET("/tasks", h.GetTasks)
	r.PUT("/tasks/:id", h.UpdateTask)
	r.DELETE("/tasks/:id", h.DeleteTask)
	return r
}

func TestCreateTask(t *testing.T) {
	router := setupRouter()

	reqBody := dto.CreateTaskRequest{
		Name:     "Test Task",
		DueDate:  nil,
		Assignee: "Tester",
		Tags:     []string{"test"},
	}
	jsonValue, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Test Task")
}

func TestGetTasks(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Task")
}

func TestGetTaskByID(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks?id=1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Task")
}

func TestCreateTask_MissingName(t *testing.T) {
	router := setupRouter()

	reqBody := dto.CreateTaskRequest{
		// Name omitted to trigger "required" validation
		Assignee: "Tester",
		Tags:     []string{"test"},
	}
	jsonValue, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Name")
}

func TestCreateTask_NameTooLong(t *testing.T) {
	router := setupRouter()

	// 101 characters to exceed max=100 rule
	longName := strings.Repeat("a", 101)
	reqBody := dto.CreateTaskRequest{
		Name:     longName,
		Assignee: "Tester",
		Tags:     []string{"test"},
	}
	jsonValue, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTaskByID_NotFound(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks?id=999", nil) // id=999 不存在
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "task not found")
}

func TestGetTaskByID_InvalidID(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks?id=abc", nil) // 非數字
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid id format")
}

func TestUpdateTask_Success(t *testing.T) {
	router := setupRouter()

	updateBody := `{"name":"Updated Task"}`
	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBufferString(updateBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestUpdateTask_InvalidID(t *testing.T) {
	router := setupRouter()

	updateBody := `{"name":"Updated Task"}`
	req, _ := http.NewRequest("PUT", "/tasks/abc", bytes.NewBufferString(updateBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteTask_Success(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteTask_NotFound(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/tasks/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteTask_InvalidID(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/tasks/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
