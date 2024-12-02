package config

import (
	"os"
)

type Config struct {
	DBName     string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
}

// NewConfig initializes new Config struct and returns pointer
func NewConfig() *Config {
	return &Config{
		DBName:     os.Getenv("MYSQL_DB"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
}
