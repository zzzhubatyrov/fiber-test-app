package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=tigserf dbname=fibertest port=5432 sslmode=disable TimeZone=Asia/Yekaterinburg"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
