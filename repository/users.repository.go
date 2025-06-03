package repository

import (
	"github.com/santigorbe/rest_golang/models"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id string) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(id string, user models.User) (models.User, error)
	Delete(id string) error
	GetByEmail(email string) (models.User, error)
	GetByEmailAndPassword(email, password string) (models.User, error)
}
