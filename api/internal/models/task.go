package models

import "time"

type Task struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Category    Category  `gorm:"foreignKey:CategoryID;references:ID"`
	Priority    string    `gorm:"type:enum('highest', 'high', 'medium', 'low', 'lowest');default:'medium';not null)"`
	Status      string    `gorm:"type:enum('pending', 'in_progress', 'completed', 'on_hold', 'blocked', 'cancelled');default:'pending';not null"`
	DueDate     time.Time `gorm:"type:timestamp"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
