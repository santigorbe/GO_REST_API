package services

import (
	"strconv"

	"github.com/santigorbe/rest_golang/models"
	"github.com/santigorbe/rest_golang/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) GetUser(id string) (models.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.Repo.Create(user)
}
func (s *UserService) UpdateUser(id string, user models.User) (models.User, error) {
	existingUser, err := s.Repo.GetByID(id)
	if err != nil {
		return models.User{}, err
	}

	user.ID = existingUser.ID // Ensure we update the correct user
	return s.Repo.Update(id, user)
}
func (s *UserService) DeleteUser(id string) error {
	existingUser2, err := s.Repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.Repo.Delete(strconv.FormatUint(uint64(existingUser2.ID), 10))
}
func (s *UserService) GetUserByEmail(email string) (models.User, error) {
	existingUser, err := s.Repo.GetByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	return existingUser, nil
}

func (s *UserService) GetUserByEmailAndPassword(email, password string) (models.User, error) {
	existingUser, err := s.Repo.GetByEmailAndPassword(email, password)
	if err != nil {
		return models.User{}, err
	}

	return existingUser, nil
}
