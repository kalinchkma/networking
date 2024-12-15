package main

import (
	"database/sql"
	"encoding/json"
	"gnja_server/internal/database"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// Request body
type requestBody struct {
	Body string `json:"body"`
}

// Error response
type errorResponse struct {
	Error string `json:"error"`
}

func createNewChirps(w http.ResponseWriter, r *http.Request) {
	// Request body
	var body requestBody

	// Body decoder
	decoder := json.NewDecoder(r.Body)

	// Decode request body
	err := decoder.Decode(&body)
	// If body decode error return error response
	if err != nil {
		// Error body
		errorRes := errorResponse{
			Error: "Client body error",
		}
		// Encode error body
		errorBody, err := json.Marshal(errorRes)

		// If encoding error return server error
		if err != nil {
			log.Printf("Error: %s", err)
			w.WriteHeader(500)
			w.Write([]byte("Internal server error"))
			return
		}
		// Return client body error
		// Add content type
		w.Header().Add("Content-Type", "application/json")
		// Add status code
		w.WriteHeader(400)
		// Attach error response body
		w.Write(errorBody)
		return
	}

	// Find the user
	users, err := cfg.GetUser(r.Context())
	// If user query error return error response
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	// If user not found return error
	if len(users) <= 0 {
		log.Printf("user not found: %v", users)
		w.WriteHeader(500)
		w.Write([]byte("There is no user to create chirps"))
		return
	}

	// Create chirps for first user: just for now
	newChirps := database.CreateCirpsParams{
		Body: sql.NullString{
			String: body.Body,
			Valid:  body.Body != "",
		},
		UserID: uuid.NullUUID{
			UUID:  users[0].ID,
			Valid: true,
		},
	}
	// Save chrip to database
	chirp, err := cfg.CreateCirps(r.Context(), newChirps)

	// Return error creating and saving is error
	if err != nil {
		log.Printf("Error: %s", err)
		w.WriteHeader(500)
		w.Write([]byte("internal server error"))
		return
	}

	// Make json string
	chripString, err := json.Marshal(chirp)

	// Return error creating json string
	if err != nil {
		log.Printf("Error: %s", err)
		w.WriteHeader(500)
		w.Write([]byte("Error marsaling response"))
		return
	}

	w.WriteHeader(201)
	w.Write(chripString)
}

func getChirps(w http.ResponseWriter, r *http.Request) {
	chirs, err := cfg.GetChirps(r.Context())

	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	// Body prepare
	resBody, err := json.Marshal(chirs)
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(resBody)
}
