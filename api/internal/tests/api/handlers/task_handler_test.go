package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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
func TestGetTaskByID(t *testing.T) {
	mockRepo := mocks.NewMockTaskRepository()
	handler := handlers.NewTaskHandler(mockRepo)

	mockTask := &models.Task{Title: "GetThisTask"}
	createdTask, _ := mockRepo.CreateTask(mockTask)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = []gin.Param{{Key: "taskID", Value: "0"}}

	handler.GetTaskByID(c)

	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code %d, got %d", http.StatusOK, recorder.Code)

	var task models.Task
	err := json.Unmarshal(recorder.Body.Bytes(), &task)
	assert.NoError(t, err, "Failed to unmarshal response body")

	assert.Equal(t, createdTask.Title, task.Title, "Expected title %s, got %s", createdTask.Title, task.Title)
	assert.Equal(t, createdTask.ID, task.ID, "Expected ID %d, got %d", createdTask.ID, task.ID)
}

func TestCreateTask(t *testing.T) {
	mockRepo := mocks.NewMockTaskRepository()
	handler := handlers.NewTaskHandler(mockRepo)

	mockTask := &models.Task{Title: "CreateTask"}
	jsonTask, _ := json.Marshal(mockTask)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("POST", "/tasks", strings.NewReader(string(jsonTask)))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.CreateTask(c)
	assert.Equal(t, http.StatusCreated, recorder.Code)

	var createdTask models.Task
	err := json.Unmarshal(recorder.Body.Bytes(), &createdTask)
	assert.NoError(t, err, "Failed to unmarshal response body")

	assert.Equal(t, mockTask.Title, createdTask.Title, "Expected title %s, got %s", mockTask.Title, createdTask.Title)
	assert.Equal(t, createdTask.ID, 0, "Expected ID 0")
}

func TestUpdateTaskByID(t *testing.T) {
	mockRepo := mocks.NewMockTaskRepository()
	handler := handlers.NewTaskHandler(mockRepo)

	mockTask := &models.Task{Title: "Original"}
	mockRepo.CreateTask(mockTask)
	assert.Equal(t, mockTask.ID, 0)

	mockTaskChanged := &models.Task{Title: "Changed"}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = []gin.Param{{Key: "taskID", Value: "0"}}

	jsonTask, err := json.Marshal(mockTaskChanged)
	assert.NoError(t, err, "error marshalling mockTask")

	c.Request = httptest.NewRequest("PUT", "/tasks/0", strings.NewReader(string(jsonTask)))
	c.Request.Header.Set("Content-Type", "application")

	handler.UpdateTaskByID(c)
	var updatedTask models.Task
	err = json.Unmarshal(recorder.Body.Bytes(), &updatedTask)
	assert.NoError(t, err, "error unmarshal updatedTask")

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, updatedTask.ID, mockTask.ID, "Expected ID %s, got %s", mockTaskChanged.ID, updatedTask.ID)
	assert.Equal(t, mockTaskChanged.Title, updatedTask.Title, "Expected title %s, got %s", mockTaskChanged.Title, updatedTask.Title)
}

func TestDeleteTaskByID(t *testing.T) {
	mockRepo := mocks.NewMockTaskRepository()
	handler := handlers.NewTaskHandler(mockRepo)

	mockTask := &models.Task{Title: "ToDelete"}
	mockRepo.CreateTask(mockTask)
	assert.Equal(t, mockTask.ID, 0)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = []gin.Param{{Key: "taskID", Value: "0"}}

	handler.DeleteTaskByID(c)
	assert.Equal(t, http.StatusNoContent, recorder.Code, "expected status code %d, got %d", http.StatusNoContent, recorder.Code)

	_, err := mockRepo.GetTaskByID(0)
	assert.Error(t, err, "Returned the element which should have been deleted")
}
