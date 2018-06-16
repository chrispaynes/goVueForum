package handlers

import (
	"net/http"
)

// GetHealth serves as a simple server health check
func GetHealth(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
