package orm

import (
	"task-api/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("task.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		panic("failed to migrate database")
	}
	return db
}
