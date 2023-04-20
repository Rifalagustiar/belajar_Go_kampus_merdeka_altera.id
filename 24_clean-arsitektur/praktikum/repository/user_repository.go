package repository

import (
	"crud/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(data models.User) error
	GetAllUsers() ([]models.User, error)
	LoginUser(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(data models.User) error {
	return r.db.Create(&data).Error
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	users := make([]models.User, 0)
	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepository) LoginUser(data models.User) (models.User, error) {
	var user models.User
	var e error
	if e = r.db.Model(&user).Where("email = ? AND password = ?", data.Email, data.Password).First(&user).Error; e != nil {
		return user, e
	}
	return user, nil
}
