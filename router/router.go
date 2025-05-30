package router

import (
	"task-api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *handler.TaskHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/tasks", handler.CreateTask)

	return r
}
