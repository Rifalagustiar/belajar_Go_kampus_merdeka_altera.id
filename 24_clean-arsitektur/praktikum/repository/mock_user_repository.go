package repository

import (
	"crud/models"

	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) Create(data models.User) error {
	ret := m.Called(data)

	return ret.Error(0)
}

func (m *mockUserRepository) GetAllUsers() ([]models.User, error) {
	ret := m.Called()

	return ret.Get(0).([]models.User), ret.Error(1)
}

func (m *mockUserRepository) LoginUser(user models.User) (models.User, error) {
	ret := m.Called(user)

	return ret.Get(0).(models.User), ret.Error(1)
}
