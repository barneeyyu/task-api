package router

import (
	"task-api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *handler.TaskHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks", handler.GetTasks)
	r.PUT("/tasks/:id", handler.UpdateTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)

	return r
}
