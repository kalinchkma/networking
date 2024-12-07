package main

import "net/http"

func handerRediness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plan; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
