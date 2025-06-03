package services

import (
	"github.com/santigorbe/rest_golang/models"
	"github.com/santigorbe/rest_golang/repository"
)

type TaskService struct {
	Repo repository.TaskRepository
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.Repo.GetAll()
}
func (s *TaskService) GetTask(id string) (models.Task, error) {
	return s.Repo.GetByID(id)
}
func (s *TaskService) CreateTask(task models.Task) (models.Task, error) {
	return s.Repo.Create(task)
}
func (s *TaskService) UpdateTask(id string, task models.Task) (models.Task, error) {
	existingTask, err := s.Repo.GetByID(id)
	if err != nil {
		return models.Task{}, err
	}

	task.ID = existingTask.ID // Ensure we update the correct task
	return s.Repo.Update(id, task)
}
func (s *TaskService) DeleteTask(id string) error {
	_, err := s.Repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.Repo.Delete(id)
}

func (s *TaskService) GetFilteredTasks(limit, offset int, search string) ([]models.Task, int64, error) {
	return s.Repo.GetFiltered(limit, offset, search)
}
