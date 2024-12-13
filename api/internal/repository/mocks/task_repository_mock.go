package mocks

import (
	"errors"

	"github.com/Cosmos307/todo-app/api/internal/models"
	"github.com/Cosmos307/todo-app/api/internal/repository"
)

type MockTaskRepository struct {
	tasks map[int]*models.Task
}

func NewMockTaskRepository() repository.TaskRepository {
	return &MockTaskRepository{
		tasks: make(map[int]*models.Task),
	}
}

func (m *MockTaskRepository) GetTasksByUserID(userID int) []models.Task {
	var tasks []models.Task
	for _, task := range m.tasks {
		if task.User.ID == userID {
			tasks = append(tasks, *task)
		}
	}
	return tasks
}

func (m *MockTaskRepository) GetTaskByID(taskID int) (*models.Task, error) {
	if task, exists := m.tasks[taskID]; exists && m.tasks[taskID].ID == taskID {
		return task, nil
	}
	return nil, errors.New("task not found")
}

func (m *MockTaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	task.ID = len(m.tasks)
	m.tasks[task.ID] = task
	return task, nil
}

func (m *MockTaskRepository) UpdateTaskByID(task *models.Task) (*models.Task, error) {
	if _, exists := m.tasks[task.ID]; exists {
		m.tasks[task.ID] = task
		return task, nil
	}
	return nil, errors.New("task not found")
}

func (m *MockTaskRepository) DeleteTaskByID(taskID int) error {
	if _, exists := m.tasks[taskID]; exists {
		delete(m.tasks, taskID)
		return nil
	}
	return errors.New("task not found")
}
