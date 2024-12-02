package models

type User struct {
	ID    int    `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(255);not null"`
	Email string `gorm:"type:varchar(255);not null; unique"`
}
