package repository

import (
	"github.com/Cosmos307/todo-app/api/internal/models"
	"gorm.io/gorm"
)

type TaskRepositoryMySQL struct {
	db *gorm.DB
}

func NewTaskRepositoryMySQL(db *gorm.DB) TaskRepository {
	return &TaskRepositoryMySQL{db: db}
}

func (r *TaskRepositoryMySQL) GetTasksByUserID(userID int) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepositoryMySQL) GetTaskByID(id int) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *TaskRepositoryMySQL) CreateTask(task *models.Task) (*models.Task, error) {
	err := r.db.Create(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepositoryMySQL) UpdateTaskByID(task *models.Task) (*models.Task, error) {
	err := r.db.Save(task).Error
	if err != nil {
		return nil, err
	}
	return task, err
}

func (r *TaskRepositoryMySQL) DeleteTaskByID(id int) error {
	return r.db.Delete(&models.Task{}, id).Error
}
