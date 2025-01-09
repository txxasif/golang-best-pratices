package services

import (
	"errors"
	"todo-api/internal/models"
)

// GetAllUsers retrieves all users from the database.
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := models.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID retrieves a user by their ID.
func GetUserByID(id uint) (models.User, error) {
	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

// GetUserByEmail retrieves a user by their email address.
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

// CreateUser creates a new user in the database.
// It validates if the email already exists before creating a new user.
func CreateUser(user models.User) (models.User, error) {
	// Check if email already exists
	existingUser, _ := GetUserByEmail(user.Email)
	if existingUser.ID != 0 {
		return user, errors.New("email already exists")
	}

	if err := models.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// UpdateUser updates the user information in the database.
func UpdateUser(user models.User) error {
	if err := models.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
