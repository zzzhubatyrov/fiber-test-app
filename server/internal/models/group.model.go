package models

import (
	"fiber-test-app/internal/config"
	"log"
	"time"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	ID        uint           `json:"ID" gorm:"primary_key"`
	Name      string         `json:"Name" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time      `json:"CreatedAt"`
	UpdatedAt time.Time      `json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `json:"DeletedAt,omitempty"`
	Todos     []Todo         `json:"Todos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func init() {
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	migrator := db.Migrator()
	if !migrator.HasTable(&Group{}) {
		if err := db.AutoMigrate(&Group{}); err != nil {
			log.Fatal(err)
			return
		}
	} else {
		if err := migrator.DropTable(&Group{}); err != nil {
			log.Fatal(err)
			return
		}
		if err := db.AutoMigrate(&Group{}); err != nil {
			log.Fatal(err)
			return
		}
	}
}
