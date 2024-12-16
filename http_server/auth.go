package main

import (
	"database/sql"
	"encoding/json"
	"gnja_server/internal/auth"
	"gnja_server/internal/database"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

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
	user, err := cfg.DB.GetUserByEmail(r.Context(), body.Email)
	// Check error
	if err != nil {
		log.Printf("Error parsing user: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("User not found"))
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

	// Generate token
	token, err := auth.MakeJWT(user.ID, cfg.JWT_SECRET, time.Second*5)
	// Check error
	if err != nil {
		log.Printf("Error creating jwt token: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	// Create the refresh token
	refreshTokenString, err := auth.MakeRefreshToken()
	// Check error
	if err != nil {
		log.Printf("Error creating refresh token: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}
	// Create new refresh token
	refreshTokenObject := database.CreateRefreshTokenParams{
		Token:     refreshTokenString,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(time.Hour * 500),
		RevokedAt: sql.NullTime{},
	}

	// Save newly created refresh token into database
	refreshToken, err := cfg.DB.CreateRefreshToken(r.Context(), refreshTokenObject)
	// Check error
	if err != nil {
		log.Printf("Error saving refresh token: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	// Success response data type
	type ResponseData struct {
		ID           uuid.UUID `json:"id"`
		CreatedAt    string    `json:"created_at"`
		UpdatedAt    string    `json:"updated_at"`
		Email        string    `json:"email"`
		Token        string    `json:"token"`
		RefreshToken string    `json:"refresh_teken"`
	}

	// Create reponse data
	responseData := ResponseData{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt.Time.String(),
		UpdatedAt:    user.UpdatedAt.Time.String(),
		Email:        user.Email,
		Token:        token,
		RefreshToken: refreshToken.Token,
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
