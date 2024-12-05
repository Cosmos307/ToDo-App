package repository

import "github.com/Cosmos307/todo-app/api/internal/models"

type CategoryRepository interface {
	GetCategoryByID(id int) (*models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategoryByID(id int) error
}
