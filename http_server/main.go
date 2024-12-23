package main

import (
	"database/sql"
	"gnja_server/internal/database"
	"gnja_server/internal/interfaces"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var cfg *interfaces.Configuration

func main() {
	const port = "8080"
	const filepathRoot = "./public"

	godotenv.Load()
	// Database connection URL
	dbURL := os.Getenv("DB_URL")

	// Open sql connection
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Printf("Error connection database: %s\n", err)
		return
	}

	cfg = &interfaces.Configuration{
		DB:         database.New(db),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}

	apiCfg := apiConfig{
		fileServerHits: atomic.Int32{},
	}

	mux := http.NewServeMux()

	// File serving
	mux.Handle("GET /app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot)))))

	// Authentication
	mux.HandleFunc("POST /api/login", middlewareLog(login))

	// Refresh token
	mux.HandleFunc("POST /api/refresh", middlewareLog(refresh))

	// Revoke token
	mux.HandleFunc("POST /api/revoke", middlewareLog(revoke))

	// Health check
	mux.HandleFunc("GET /api/healthz", middlewareLog(handerRediness))

	// Validate json body
	mux.HandleFunc("POST /api/validate_zingy", middlewareLog(validate_zingy))

	// Create user
	mux.HandleFunc("POST /api/users", middlewareLog(createNewUser))

	// Update user
	mux.HandleFunc("PUT /api/users", middlewareLog(updateUsers))

	// Create chirp
	mux.HandleFunc("POST /api/chirps", middlewareLog(createNewChirps))

	// Get all chirps
	mux.HandleFunc("GET /api/chirps", middlewareLog(getChirps))

	// Get chirps
	mux.HandleFunc("GET /api/chirps/{chirpID}", middlewareLog(getChirpsByID))

	// Metrics
	mux.HandleFunc("GET /admin/metrics", middlewareLog(apiCfg.handlerMetrics))

	// Reset counter
	mux.HandleFunc("POST /admin/reset", middlewareLog(apiCfg.handerReset))

	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)

	log.Fatal(server.ListenAndServe())
}
