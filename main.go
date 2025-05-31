package main

import (
	"task-api/handler"
	"task-api/pkg/orm"
	"task-api/repository"
	"task-api/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "task-api/docs" // 🔥 非常重要！引用剛剛產生的 Swagger 文件
)

// @title Task API
// @version 1.0
// @description A simple RESTful API for managing tasks.
// @host localhost:8080
// @BasePath /
func main() {
	db := orm.InitDB()
	repo := repository.NewTaskRepository(db)
	handler := handler.NewTaskHandler(repo)

	r := router.SetupRouter(handler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080") // 啟動 server
}
