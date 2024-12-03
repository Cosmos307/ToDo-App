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

func (r *TaskRepositoryMySQL) CreateTask(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepositoryMySQL) UpdateTask(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepositoryMySQL) DeleteTaskByID(id int) error {
	return r.db.Delete(&models.Task{}, id).Error
}
