package usecase

import (
	"crud/models"
	"crud/repository"
)

type UserUsecase interface {
	Create(payload models.User) error
	GetAllUsers() ([]models.User, error)
	LoginUser(payload models.User) (models.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (s *userUsecase) Create(payload models.User) error {
	err := s.userRepository.Create(payload)

	if err != nil {
		return err
	}

	return nil
}

func (s *userUsecase) GetAllUsers() ([]models.User, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		return users, err

	}
	return users, nil
}

func (s *userUsecase) LoginUser(payload models.User) (models.User, error) {
	user, err := s.userRepository.LoginUser(payload)
	if err != nil {
		return user, err
	}
	return user, nil
}
