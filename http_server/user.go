package main

import (
	"encoding/json"
	"gnja_server/internal/auth"
	"gnja_server/internal/database"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func createNewUser(w http.ResponseWriter, r *http.Request) {
	// Request body
	type RequestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	reqBody := RequestBody{}
	err := decoder.Decode(&reqBody)
	if err != nil {
		log.Printf("Error parsing request body: %s", err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Error input"))
		return
	}

	// hash password
	HashedPassword, err := auth.HashPassword(reqBody.Password)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	newUser := database.CreateUserParams{
		Email:          reqBody.Email,
		HashedPassword: HashedPassword,
	}

	// save the user
	user, err := cfg.CreateUser(r.Context(), newUser)

	if err != nil {
		log.Printf("Error saving user to database: %s", err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	resBody, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error: %s", err)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(resBody)
}

func login(w http.ResponseWriter, r *http.Request) {
	// Request body
	type requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// json decoder
	decoder := json.NewDecoder(r.Body)
	body := requestBody{}
	err := decoder.Decode(&body)
	if err != nil {
		log.Printf("body decoing error: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	// Find the user by email
	user, err := cfg.GetUserByEmail(r.Context(), body.Email)
	// Check error
	if err != nil {
		log.Printf("Error parsing user: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	// Compare password
	err = auth.CheckPasswordHash(body.Password, user.HashedPassword)
	// Check error
	if err != nil {
		log.Printf("Error comparing password: %v\n", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	type ResponseData struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt string    `json:"created_at"`
		UpdatedAt string    `json:"updated_at"`
		Email     string    `json:"email"`
	}

	responseData := ResponseData{
		ID:        user.ID,
		CreatedAt: user.CreatedAt.Time.String(),
		UpdatedAt: user.UpdatedAt.Time.String(),
		Email:     user.Email,
	}

	// Prepare login success object
	resObject, err := json.Marshal(responseData)
	// Check error
	if err != nil {
		log.Printf("Error marshaling user object: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	// Return success if passowrd matched
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(resObject)
}
