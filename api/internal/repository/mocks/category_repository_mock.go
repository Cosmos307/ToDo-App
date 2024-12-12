package mocks

import (
	"errors"

	"github.com/Cosmos307/todo-app/api/internal/models"
	"github.com/Cosmos307/todo-app/api/internal/repository"
)

type MockCategoryRepository struct {
	categories map[int]*models.Category
}

func NewMockCategoryRepository() repository.CategoryRepository {
	return &MockCategoryRepository{
		categories: make(map[int]*models.Category),
	}
}

func (m *MockCategoryRepository) GetCategoryByID(categoryID int) (*models.Category, error) {
	if category, exists := m.categories[categoryID]; exists {
		return category, nil
	}
	return nil, errors.New("category not found")
}

func (m *MockCategoryRepository) CreateCategory(category *models.Category) (*models.Category, error) {
	category.ID = len(m.categories)
	m.categories[category.ID] = category
	return m.categories[category.ID], nil
}

func (m *MockCategoryRepository) UpdateCategoryByID(category *models.Category) (*models.Category, error) {
	if _, exists := m.categories[category.ID]; exists {
		m.categories[category.ID] = category
		return category, nil
	}
	return nil, errors.New("category not found")
}

func (m *MockCategoryRepository) DeleteCategoryByID(categoryID int) error {
	if _, exists := m.categories[categoryID]; exists {
		delete(m.categories, categoryID)
		return nil
	}
	return errors.New("category not found")
}
