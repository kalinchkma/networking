package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

func main() {
	const port = "8080"
	const filepathRoot = "./public"

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
