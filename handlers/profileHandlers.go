package handlers

import (
	"encoding/json"
	"fmt"
	"go-todo/dtos"
	"go-todo/models"
	"math/rand/v2"
	"net/http"
)

func CreateProfileHandler(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateProfileRequestDTO

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Create a JSON decoder and disallow unknown fields
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Decode the JSON body into the dto struct
	err := decoder.Decode(&dto)

	if err != nil {
		http.Error(w, "Invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}

	validEmail, err := models.NewEmail(dto.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validPassword, err := models.NewPassword(dto.Password, dto.RepeatPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	println(validEmail, validPassword)

	// Your business logic here
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "profile/%d", rand.Uint32())
}
