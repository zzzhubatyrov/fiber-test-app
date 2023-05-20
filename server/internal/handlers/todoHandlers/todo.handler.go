package todoHandlers

import (
	"fiber-test-app/internal/config"
	"fiber-test-app/internal/models"
)

func CheckTodo(model []models.Todo) ([]models.Todo, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	if err := db.Find(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func CreateTodo(models []models.Todo) ([]models.Todo, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	if err := db.Create(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func DeleteTodo(id uint) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	var todo models.Todo
	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return err
	}
	if err := db.Delete(&todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(id string, todo *models.Todo) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}
	var existingTodo models.Todo
	if err := db.Where("id = ?", id).First(&existingTodo).Error; err != nil {
		return err
	}

	existingTodo.Title = todo.Title
	existingTodo.Description = todo.Description
	existingTodo.Completed = todo.Completed

	if err := db.Save(&existingTodo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodo(id string) (*models.Todo, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}

	var todo models.Todo
	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}
