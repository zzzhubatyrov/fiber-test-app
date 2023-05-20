package models

import (
	"fiber-test-app/internal/config"
	"log"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID          uint           `json:"ID" gorm:"primary_key"`
	Title       string         `json:"Title" gorm:"type:varchar(100);not null"`
	Description string         `json:"Description" gorm:"type:varchar(255);not null"`
	Completed   bool           `json:"Completed" gorm:"type:boolean;default:false"`
	CreatedAt   time.Time      `json:"CreatedAt"`
	UpdatedAt   time.Time      `json:"UpdatedAt"`
	DeletedAt   gorm.DeletedAt `json:"DeletedAt,omitempty"`
}

func init() {
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	migrator := db.Migrator()
	if !migrator.HasTable(Todo{}) {
		if err := db.AutoMigrate(&Todo{}); err != nil {
			log.Fatal(err)
			return
		}
	} else {
		if err := migrator.DropTable(&Todo{}); err != nil {
			log.Fatal(err)
			return
		}
		if err := db.AutoMigrate(&Todo{}); err != nil {
			log.Fatal(err)
			return
		}
	}
}
