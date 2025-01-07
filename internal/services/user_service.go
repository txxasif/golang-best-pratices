package services

import (
	"errors"
	"todo-api/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := models.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	if err := models.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
