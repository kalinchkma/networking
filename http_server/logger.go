package main

import (
	"log"
	"net/http"
)

func middlewareLog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
