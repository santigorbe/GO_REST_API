package repository

import (
	"errors"
	"strconv"

	"github.com/santigorbe/rest_golang/db"
	"github.com/santigorbe/rest_golang/models"
	"gorm.io/gorm"
)

type GormTaskRepository struct{}

func (r *GormTaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	result := db.DB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *GormTaskRepository) GetByID(id string) (models.Task, error) {
	var task models.Task

	result := db.DB.First(&task, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Task{}, errors.New("task not found")
	}
	if result.Error != nil {
		return models.Task{}, result.Error
	}

	return task, nil
}

func (r *GormTaskRepository) Create(task models.Task) (models.Task, error) {
	result := db.DB.Create(&task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return task, nil
}

func (r *GormTaskRepository) Update(id string, task models.Task) (models.Task, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}
	task.ID = uint(intID)
	result := db.DB.Save(&task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return task, nil
}

func (r *GormTaskRepository) Delete(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("task not found")
	}
	result := db.DB.Delete(&models.Task{}, intID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (r *GormTaskRepository) GetByUserID(userID string) ([]models.Task, error) {
	var tasks []models.Task
	result := db.DB.Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *GormTaskRepository) GetFiltered(limit, offset int, search string) ([]models.Task, int64, error) {
	var tasks []models.Task
	var total int64

	query := db.DB.Model(&models.Task{})

	if search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Limit(limit).Offset(offset).Find(&tasks).Error
	return tasks, total, err
}
