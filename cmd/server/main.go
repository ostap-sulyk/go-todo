package main

import (
	"context"
	"go-todo/internal/database"
	"go-todo/internal/handlers"
	"go-todo/internal/repository" // Import the sqlc-generated package
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {


	ctx := context.Background()
	// Database connection
	pool, err := database.NewPool(ctx)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer pool.Close()

	// Initialize repository with sqlc-generated Queries
	queries := repository.New(pool)

	// Initialize handler
	userHandler := handlers.NewUserHandler(queries)

	// Routes
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),
	)

	r.Post("/users/create", userHandler.CreateUserHandler)

	// Start the server
	log.Println("Server started on :3000")
	http.ListenAndServe(":3000", r)
}
