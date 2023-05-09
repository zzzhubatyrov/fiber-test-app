package todoHandler

import (
	"fiber-test-app/internal/config"
	"fiber-test-app/internal/models"
)

func CreateModel(model []models.Todo) ([]models.Todo, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model)
	return model, nil
}
