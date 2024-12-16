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

func updateUsers(w http.ResponseWriter, r *http.Request) {
	// Request body
	type RequestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Get authorized token
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("Error parsing authorized token: %v", err)
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized User"))
		return
	}

	// Validate user token
	userID, err := auth.ValidateJWT(token, cfg.JWT_SECRET)
	if err != nil {
		log.Printf("Error validating user token: %v", err)
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized user"))
		return
	}

	// Find user by id
	_, err = cfg.DB.GetUserByID(r.Context(), userID)
	if err != nil {
		log.Printf("Error user not found: %v", err)
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized user"))
		return
	}

	// Parse user request input
	decoder := json.NewDecoder(r.Body)
	body := &RequestBody{}
	err = decoder.Decode(body)
	if err != nil {
		log.Printf("Error parsing user inputs: %v", err)
		w.WriteHeader(400)
		w.Write([]byte("There is an error of your inputs"))
		return
	}

	// Hash new password
	hashedPassowrd, err := auth.HashPassword(body.Password)
	if err != nil {
		log.Printf("Error creating hash password: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	// Prepare for user update
	updateUserData := database.UpdateUserByIDParams{
		Email:          body.Email,
		HashedPassword: hashedPassowrd,
		ID:             userID,
	}
	// Update the user
	err = cfg.DB.UpdateUserByID(r.Context(), updateUserData)
	if err != nil {
		log.Printf("Error updating the user: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User updated successfully"))

}
