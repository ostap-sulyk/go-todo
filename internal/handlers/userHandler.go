package handlers

import (
	"encoding/json"
	"fmt"
	"go-todo/internal/models"
	"go-todo/internal/repository"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repo *repository.Queries
}

func NewUserHandler(repo *repository.Queries) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var createUserRequest models.CreateUserRequest

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&createUserRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate fields
	if createUserRequest.Email == "" || createUserRequest.Password == "" || createUserRequest.RepeatPassword == "" {
		http.Error(w, "Email, password, and repeat password are required", http.StatusBadRequest)
		return
	}

	// Check if passwords match
	if createUserRequest.Password != createUserRequest.RepeatPassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	params := repository.CreateUserParams{
		Email:    createUserRequest.Email,
		Password: string(hashedPassword),
	}

	// Insert the user into the database
	user, err := h.repo.CreateUser(r.Context(), params)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create user: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
