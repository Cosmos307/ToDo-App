package database

import (
	"fmt"
	"log"
	"time"

	"github.com/Cosmos307/todo-app/api/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	// Set Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var db *gorm.DB
	var err error
	maxRetries := 5
	retryInterval := 5 * time.Second

	// Connect to the database with retries
	for i := 0; i < maxRetries; i++ {
		// Connect to the database
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Error connecting to the database: %v", err)
			time.Sleep(5 * retryInterval)
			continue
		}

		// Extract the underlying *sql.DB instance to ping test and ensure a valid connection
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting database instance: %v", err)
			time.Sleep(retryInterval)
			continue
		}
		if pingError := sqlDB.Ping(); pingError != nil {
			log.Printf("Error pinging the database: %v", pingError)
			time.Sleep(5 * retryInterval)
			continue
		}

		fmt.Println("Success connecting to the database")
		return db, nil
	}

	return nil, fmt.Errorf("failed to connect to the database after %d attempts: %v", maxRetries, err)
}
