package api

import "net/http"

// HomeHandler is the handler for the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to API Gateway"))
}
