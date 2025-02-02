package repository

import (
	"github.com/Cosmos307/todo-app/api/internal/models"
	"gorm.io/gorm"
)

type CategoryRepositoryMySQL struct {
	db *gorm.DB
}

func NewCategoryRepositoryMySQL(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryMySQL{db: db}
}

func (r *CategoryRepositoryMySQL) GetCategoryByID(categoryID int) (*models.Category, error) {
	var category *models.Category
	err := r.db.First(category, categoryID).Error
	return category, err
}

// check if category variable contains new id set by database
func (r *CategoryRepositoryMySQL) CreateCategory(category *models.Category) (*models.Category, error) {
	err := r.db.Create(category).Error
	if err != nil {
		return nil, err
	}
	return category, err
}

func (r *CategoryRepositoryMySQL) UpdateCategoryByID(category *models.Category) (*models.Category, error) {
	err := r.db.Save(category).Error
	if err != nil {
		return nil, err
	}
	return category, err
}

func (r *CategoryRepositoryMySQL) DeleteCategoryByID(categoryID int) error {
	err := r.db.Delete(&models.Category{}, categoryID).Error
	return err
}
