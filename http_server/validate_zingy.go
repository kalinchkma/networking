package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func validate_zingy(w http.ResponseWriter, r *http.Request) {
	// Request body
	type requestBody struct {
		Body string `json:"body" validate:"required"`
	}

	// Error respose body
	type errorResponseBody struct {
		Error string `json:"error"`
	}
	// Success response body
	type successResponseBody struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	body := requestBody{}
	err := decoder.Decode(&body)
	if err != nil {
		// error body
		errorBody := errorResponseBody{
			Error: "Something went wrong",
		}
		errorData, err := json.Marshal(errorBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			w.Write([]byte("Internal server error"))
			return
		}
		log.Printf("Error request body is not in good format: %s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(errorData)
		return
	}

	if len(body.Body) > 140 {
		errorBody := errorResponseBody{
			Error: "Zingy is too long",
		}
		errorData, err := json.Marshal(errorBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			w.Write([]byte("Internal server error"))
			return
		}
		log.Printf("Error request body is not in good format: %s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(errorData)
		return
	}

	successBody := successResponseBody{
		Valid: true,
	}

	successData, err := json.Marshal(successBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(successData)
}
