package repository

import (
	"github.com/santigorbe/rest_golang/models"
)

type TaskRepository interface {
	GetAll() ([]models.Task, error)
	GetByID(id string) (models.Task, error)
	Create(task models.Task) (models.Task, error)
	Update(id string, task models.Task) (models.Task, error)
	Delete(id string) error
	GetByUserID(userID string) ([]models.Task, error)
	GetFiltered(limit, offset int, search string) ([]models.Task, int64, error)
}
