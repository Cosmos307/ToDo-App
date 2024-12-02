package models

import "time"

type Category struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"type:varchar(255);not null"`
	User      User      `gorm:"foreignKey:User;references:ID"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
