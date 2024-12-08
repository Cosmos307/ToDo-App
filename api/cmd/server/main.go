package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Cosmos307/todo-app/api/internal/config"
	"github.com/Cosmos307/todo-app/api/internal/database"
	"github.com/Cosmos307/todo-app/api/internal/handlers"
	"github.com/Cosmos307/todo-app/api/internal/repository"
	"github.com/Cosmos307/todo-app/api/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.NewConfig()
	database, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	if database == nil {
		log.Fatal("Database connection is nil")
	}

	router := gin.Default()

	userRepository := repository.NewUserRepositoryMySQL(database)
	taskRepository := repository.NewTaskRepositoryMySQL(database)
	categoryRepository := repository.NewCategoryRepositoryMySQL(database)

	userHandler := handlers.NewUserHandler(userRepository)
	taskHandler := handlers.NewTaskHandler(taskRepository)
	categoryHandler := handlers.NewCategoryHandler(categoryRepository)

	routes.RegisterRoutes(router, userHandler, taskHandler, categoryHandler)
	port := os.Getenv("API_PORT")
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
