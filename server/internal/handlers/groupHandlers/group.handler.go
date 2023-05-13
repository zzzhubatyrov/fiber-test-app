package groupHandlers

import (
	"fiber-test-app/internal/config"
	"fiber-test-app/internal/models"
)

func CreateGroup(group models.Group) (models.Group, error) {
	db, err := config.Connect()
	if err != nil {
		return models.Group{}, err
	}
	if err := db.Create(&group).Error; err != nil {
		return models.Group{}, err
	}
	return group, nil
}
