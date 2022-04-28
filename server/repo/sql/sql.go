package sql

import (
	"log"

	"github.com/ranefattesingh/task-management-app/server/data"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sql struct {
	db *gorm.DB
}

func NewSql(dsn string, pool int) *sql {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	applyMigrations(db)

	return &sql{db}
}

func applyMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&data.Task{})
	if err != nil {
		log.Fatal(err)
	}
}
