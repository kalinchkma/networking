package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func createNewUser(w http.ResponseWriter, r *http.Request) {
	// Request body
	type RequestBody struct {
		Email string `json:"email"`
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

	// save the user
	user, err := cfg.CreateUser(r.Context(),
		sql.NullString{
			String: reqBody.Email,
			Valid:  reqBody.Email != "",
		})
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
