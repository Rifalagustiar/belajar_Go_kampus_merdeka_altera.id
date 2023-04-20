package usecase

import (
	"crud/models"
	"crud/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserCreateSuccess(t *testing.T) {
	mockUser := models.User{
		Email:    "test@gmail.com",
		Password: "12345678",
	}

	mockUserRepository := repository.NewMockUserRepository()

	mockUserRepository.On("Create", mockUser).Return(nil)

	service := NewUserUsecase(mockUserRepository)

	if err := service.Create(mockUser); err != nil {
		t.Errorf("Got Error %v", err)
	}

}

func TestUserCreateFailed(t *testing.T) {
	mockUser := models.User{
		Email:    "Rifal@gmail.com",
		Password: "12345678",
	}

	mockUserRepository := repository.NewMockUserRepository()

	mockUserRepository.On("Create", mockUser).Return(errors.New("failed"))

	service := NewUserUsecase(mockUserRepository)

	if err := service.Create(mockUser); err != nil {
		// t.Errorf("Got Error %v", err)
		assert.Error(t, err, "failed")

	}

}

func TestUserGetAllUsersSuccess(t *testing.T) {

	mockUser := make([]models.User, 0)

	mockUser = append(mockUser, models.User{
		Email:    "test1@gmail.com",
		Password: "12345678",
	})

	mockUserRepository := repository.NewMockUserRepository()
	mockUserRepository.On("GetAllUsers").Return(mockUser, nil)

	service := NewUserUsecase(mockUserRepository)

	users, err := service.GetAllUsers()

	if err != nil {
		t.Errorf("Got Error %v", err)
	}

	assert.Equal(t, users[0].Email, mockUser[0].Email)
}
func TestUserGetAllUsersFailed(t *testing.T) {

	mockUser := make([]models.User, 0)

	mockUser = append(mockUser, models.User{
		Email:    "test1@gmail.com",
		Password: "12345678",
	})

	mockUserRepository := repository.NewMockUserRepository()
	mockUserRepository.On("GetAllUsers").Return(mockUser, errors.New("failed"))

	service := NewUserUsecase(mockUserRepository)

	users, err := service.GetAllUsers()

	if err != nil {
		// t.Errorf("Got Error %v", err)
		assert.Error(t, err, "failed")
	}

	assert.Equal(t, users[0].Email, mockUser[0].Email)
}

func TestLoginUserSuccess(t *testing.T) {
	mockUser := models.User{
		Email:    "test1@gmail.com",
		Password: "12345678",
	}

	mockUserRepository := repository.NewMockUserRepository()

	mockUserRepository.On("LoginUser", mockUser).Return(mockUser, nil)

	service := NewUserUsecase(mockUserRepository)
	user, err := service.LoginUser(mockUser)
	if err != nil {
		t.Errorf("Got Error %v", err)
	}

	assert.Equal(t, user.Email, mockUser.Email)

}

func TestLoginUserFailed(t *testing.T) {
	mockUser := models.User{
		Email:    "test1@gmail.com",
		Password: "12345678",
	}

	mockUserRepository := repository.NewMockUserRepository()

	mockUserRepository.On("LoginUser", mockUser).Return(mockUser, errors.New("failed"))

	service := NewUserUsecase(mockUserRepository)
	user, err := service.LoginUser(mockUser)
	if err != nil {
		assert.Error(t, err, "failed")
	}
	assert.Equal(t, user.Email, mockUser.Email)

}
