package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestCreateCategory(t *testing.T) {
	mockRepo := mocks.NewMockCategoryRepository()
	handler := handlers.NewCategoryHandler(mockRepo)

	mockCategory := &models.Category{Title: "CreatedCategory"}
	jsonCategory, err := json.Marshal(mockCategory)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("POST", "/categories/", strings.NewReader(string(jsonCategory)))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.CreateCategory(c)
	assert.Equal(t, http.StatusCreated, recorder.Code)

	createdCategory, err := mockRepo.GetCategoryByID(0)
	assert.NoError(t, err, "Failed to get created category")
	assert.Equal(t, mockCategory.Title, createdCategory.Title, "Expected category title %s, got %s", mockCategory.Title, createdCategory.Title)
}

func TestUpdateCategoryByID(t *testing.T) {
	mockRepo := mocks.NewMockCategoryRepository()
	handler := handlers.NewCategoryHandler(mockRepo)

	mockCategory := &models.Category{Title: "Before update"}
	mockRepo.CreateCategory(mockCategory)

	mockCategory.Title = "After udpate"
	updatedCategory := mockCategory
	jsonUpdatedCategory, err := json.Marshal(&updatedCategory)
	assert.NoError(t, err)

	updateRecorder := httptest.NewRecorder()
	updateC, _ := gin.CreateTestContext(updateRecorder)
	updateC.Params = []gin.Param{{Key: "categoryID", Value: "0"}}
	updateC.Request = httptest.NewRequest("PUT", "/categories/1", strings.NewReader(string(jsonUpdatedCategory)))
	updateC.Request.Header.Set("Content-Type", "application/json")

	handler.UpdateCategoryByID(updateC)
	assert.Equal(t, http.StatusOK, updateRecorder.Code, "UpdateCategoryByID func expected status code %d, got %d", http.StatusOK, updateRecorder.Code)

	updatedCategoryFromRepo, err := mockRepo.GetCategoryByID(0)
	assert.NoError(t, err, "Failed to get updated category")
	assert.Equal(t, &updatedCategory, &updatedCategoryFromRepo, "Expected category title %s, got %s", updatedCategory.Title, updatedCategoryFromRepo.Title)
}
