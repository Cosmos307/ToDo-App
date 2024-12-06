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

	categoryRepo := repository.NewCategoryRepositoryMySQL(database)
	categoryHandler := handlers.NewCategoryHandler(categoryRepo)

	router := gin.Default()

	routes.RegisterRoutes(router, categoryHandler)

	port := os.Getenv("API_PORT")
	fmt.Printf("Server running on port %s\n", port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
