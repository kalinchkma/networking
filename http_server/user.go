package main

import (
	"encoding/json"
	"gnja_server/internal/auth"
	"gnja_server/internal/database"
	"log"
	"net/http"
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
	user, err := cfg.DB.CreateUser(r.Context(), newUser)

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
