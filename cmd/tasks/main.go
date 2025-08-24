package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
	"task-planner/internal/repository"
	"task-planner/internal/service"
	"task-planner/internal/web"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		// This is not a fatal error. It just means the .env file was not found.
		// The app can still run if the environment variables are set manually.
		log.Println("Warning: No .env file found, relying on system environment variables.")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("Error: DB_PASSWORD environment variable is not set. Please set it in a .env file or export it.")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", dbPassword, "postgres")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	log.Println("Successfully connected to PostgreSQL!")

	taskRepository := repository.NewRepository(db)
	if err := taskRepository.CreateTable(); err != nil {
		log.Fatal(err)
	}

	taskService := service.NewTaskService(taskRepository)
	taskHandler := web.NewTaskHandler(taskService)

	router := http.NewServeMux()
	router.HandleFunc("GET /v1/tasks", taskHandler.GetTasks)
	router.HandleFunc("POST /v1/tasks", taskHandler.CreateTask)
	router.HandleFunc("GET /v1/tasks/{id}", taskHandler.GetTaskByID)
	router.HandleFunc("PUT /v1/tasks/{id}", taskHandler.UpdateTask)
	router.HandleFunc("DELETE /v1/tasks/{id}", taskHandler.DeleteTask)

	log.Println("Server starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}