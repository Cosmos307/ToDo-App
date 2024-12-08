package repository

import (
	"github.com/Cosmos307/todo-app/api/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryMySQL struct {
	db *gorm.DB
}

func NewUserRepositoryMySQL(db *gorm.DB) UserRepository {
	return &UserRepositoryMySQL{db: db}
}

func (r *UserRepositoryMySQL) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, userID).Error
	return &user, err
}

func (r *UserRepositoryMySQL) CreateUser(user *models.User) error {
	err := r.db.Create(user).Error
	return err
}

func (r *UserRepositoryMySQL) UpdateUserByID(user *models.User) error {
	err := r.db.Save(user).Error
	return err
}

func (r *UserRepositoryMySQL) DeleteUserByID(userID int) error {
	err := r.db.Delete(&models.User{}, userID).Error
	return err
}
