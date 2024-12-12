package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Cosmos307/todo-app/api/internal/handlers"
	"github.com/Cosmos307/todo-app/api/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/Cosmos307/todo-app/api/internal/repository/mocks"
)

func TestGetCategoryByID(t *testing.T) {
	mockRepo := mocks.NewMockCategoryRepository()
	handler := handlers.NewCategoryHandler(mockRepo)

	mockCategory := &models.Category{ID: 1, Title: "Test Get Method"}
	_, err := mockRepo.CreateCategory(mockCategory)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = []gin.Param{{Key: "categoryID", Value: "0"}}

	handler.GetCategoryByID(c)

	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code %d, got %d", http.StatusOK, recorder.Code)
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	var category models.Category
	err = json.Unmarshal(recorder.Body.Bytes(), &category)
	assert.NoError(t, err, "Failed to unmarshal response body")

	assert.Equal(t, mockCategory.ID, category.ID, "Expected category ID %d, got %d", mockCategory.ID, category.ID)
	assert.Equal(t, mockCategory.Title, category.Title, "Expected category title %s, got %s", mockCategory.Title, category.Title)

}
