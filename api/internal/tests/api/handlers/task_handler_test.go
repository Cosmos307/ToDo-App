package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Cosmos307/todo-app/api/internal/handlers"
	"github.com/Cosmos307/todo-app/api/internal/models"
	"github.com/Cosmos307/todo-app/api/internal/repository/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTaskByUserID(t *testing.T) {

	mockRepo := mocks.NewMockTaskRepository()
	handler := handlers.NewTaskHandler(mockRepo)

	mockTasks := []models.Task{
		{Title: "title1", User: models.User{ID: 1}},
		{Title: "title2", User: models.User{ID: 1}},
	}

	mockRepo.CreateTask(&mockTasks[0])
	mockRepo.CreateTask(&mockTasks[1])

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = []gin.Param{{Key: "userID", Value: "1"}}

	handler.GetTasksByUserID(c)

	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code %d, got %d", http.StatusOK, recorder.Code)
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	var tasks []models.Task
	err := json.Unmarshal(recorder.Body.Bytes(), &tasks)
	assert.NoError(t, err, "Failed to unmarshal response body")

	assert.Equal(t, mockTasks[0], tasks[0], "Expected mockTask %d, got %d", mockTasks[0], tasks[0])
	assert.Equal(t, mockTasks[1], tasks[1], "Expected mockTask %d, got %d", mockTasks[1], tasks[1])

}

/*
type TaskHandler struct {
	repo repository.TaskRepository
}

func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func (h *TaskHandler) GetTasksByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	tasks := h.repo.GetTasksByUserID(userID)
	c.JSON(http.StatusOK, tasks)
}
*/
