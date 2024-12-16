package main

import (
	"database/sql"
	"encoding/json"
	"gnja_server/internal/auth"
	"gnja_server/internal/database"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Success response data type
type ResponseData struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
	Email        string    `json:"email"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_teken"`
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

func refresh(w http.ResponseWriter, r *http.Request) {
	// Extract the token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Authorization header must start with 'Bearer '", http.StatusUnauthorized)
		return
	}

	refreshToken := strings.TrimPrefix(authHeader, "Bearer ")

	// Validate the refresh token in the database
	dbToken, err := cfg.DB.GetRefreshTokenByToken(r.Context(), refreshToken)

	if err != nil || (dbToken.RevokedAt != sql.NullTime{}) || time.Now().After(dbToken.ExpiresAt) {
		http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	// Generate the access token
	accessToken, err := auth.MakeJWT(dbToken.UserID, cfg.JWT_SECRET, time.Second*60)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	// Response with access token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"token": ` + accessToken + `}`))
}

func revoke(w http.ResponseWriter, r *http.Request) {
	// Extract token from Autorization header
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Authorization header must start with 'Bearer '", http.StatusUnauthorized)
		return
	}

	refreshToken := strings.TrimPrefix(authHeader, "Bearer ")

	// Revoke the refresh token
	err := cfg.DB.RevokeRefreshToken(r.Context(), refreshToken)
	if err != nil {
		http.Error(w, "Failed to revoke token", http.StatusUnauthorized)
		return
	}

	// Response with 204 No context
	w.WriteHeader(http.StatusNoContent)
}
