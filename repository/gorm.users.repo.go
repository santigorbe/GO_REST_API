package repository

import (
	"errors"
	"strconv"

	"github.com/santigorbe/rest_golang/db"
	"github.com/santigorbe/rest_golang/models"
	"gorm.io/gorm"
)

type GormUserRepository struct{}

func (r *GormUserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	result := db.DB.Find(&users)
	return users, result.Error
}
func (r *GormUserRepository) GetByID(id string) (models.User, error) {
	var user models.User

	result := db.DB.Preload("Tasks").First(&user, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("user not found")
	}
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (r *GormUserRepository) Create(user models.User) (models.User, error) {
	result := db.DB.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (r *GormUserRepository) Update(id string, user models.User) (models.User, error) {

	// Check if the user exists
	if _, err := r.GetByID(id); err != nil {
		return models.User{}, err
	}

	// Update the user
	intID, err := strconv.Atoi(id)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	user.ID = uint(intID)
	result := db.DB.Save(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (r *GormUserRepository) Delete(id string) error {
	// Check if the user exists
	if _, err := r.GetByID(id); err != nil {
		return err
	}
	result := db.DB.Delete(&models.User{}, id)
	return result.Error
}

func (r *GormUserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	result := db.DB.Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("user not found")
	}
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (r *GormUserRepository) GetByEmailAndPassword(email, password string) (models.User, error) {
	var user models.User
	result := db.DB.Where("email = ? AND password = ?", email, password).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("user not found")
	}
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
