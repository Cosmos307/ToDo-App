package repository

import "github.com/Cosmos307/todo-app/api/internal/models"

type CategoryRepository interface {
	GetCategoryByID(id int) (*models.Category, error)
	CreateCategory(category *models.Category) (*models.Category, error)
	UpdateCategoryByID(category *models.Category) (*models.Category, error)
	DeleteCategoryByID(id int) error
}
