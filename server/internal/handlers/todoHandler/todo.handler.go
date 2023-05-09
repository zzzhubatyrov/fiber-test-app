package todoHandler

import (
	"fiber-test-app/internal/config"
	"fiber-test-app/internal/models"
)

func CheckTodo(model []models.Todo) ([]models.Todo, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}

	db.Find(&model)
	return model, nil
}
