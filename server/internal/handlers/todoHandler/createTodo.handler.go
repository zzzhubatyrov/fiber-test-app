package todoHandler

import (
	"fiber-test-app/internal/config"
	"fiber-test-app/internal/models"
)

func CreateTodo(model []models.Todo) ([]models.Todo, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	db.Create(&model)
	return model, nil
}
