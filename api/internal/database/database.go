package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Load database configuration data from .env and connect api with database
func Init() {
	var err error

	// Load environment variables
	dbName := os.Getenv("MYSQL_DB")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Set Data Source Name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to the database with retries
	for {
		// Connect to the database
		if db, err = sql.Open("mysql", dsn); err != nil {
			log.Fatalf("Error connecting to the database: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Test connection
		if pingError := db.Ping(); pingError != nil {
			log.Fatalf("Error pinging the database: %v", pingError)
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("Success connecting to the database")
		break
	}
}

func GetDB() *sql.DB {
	return db
}
