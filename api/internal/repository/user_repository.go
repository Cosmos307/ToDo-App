package repository

import "github.com/Cosmos307/todo-app/api/internal/models"

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUserByID(id int) error
}
