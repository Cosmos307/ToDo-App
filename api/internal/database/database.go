package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Cosmos307/todo-app/api/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	// Set Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var db *sql.DB
	var err error
	maxRetries := 5
	retryInterval := 5 * time.Second

	// Connect to the database with retries
	for i := 0; i < maxRetries; i++ {
		// Connect to the database
		if db, err = sql.Open("mysql", dsn); err != nil {
			log.Printf("Error connecting to the database: %v", err)
			time.Sleep(5 * retryInterval)
			continue
		}

		// Test connection
		if pingError := db.Ping(); pingError != nil {
			log.Printf("Error pinging the database: %v", pingError)
			time.Sleep(5 * retryInterval)
			continue
		}

		fmt.Println("Success connecting to the database")
		return db, nil
	}

	return nil, fmt.Errorf("failed to connect to the database after %d attempts: %v", maxRetries, err)
}
