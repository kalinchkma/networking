package main

import (
	"database/sql"
	"gnja_server/internal/database"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var cfg *database.Queries

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

	cfg = database.New(db)

	apiCfg := apiConfig{
		fileServerHits: atomic.Int32{},
	}

	mux := http.NewServeMux()

	// File serving
	mux.Handle("GET /app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot)))))

	// Health check
	mux.HandleFunc("GET /api/healthz", middlewareLog(handerRediness))

	// Validate json body
	mux.HandleFunc("POST /api/validate_zingy", middlewareLog(validate_zingy))

	// Create user
	mux.HandleFunc("POST /api/users", middlewareLog(createNewUser))

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
