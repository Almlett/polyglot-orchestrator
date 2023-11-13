package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "PolyglotOrchestrator/internal/api"
)

func main() {
    r := mux.NewRouter()

    // Routes consist of a path and a handler function.
    r.HandleFunc("/", api.HomeHandler)

    // Start server
    log.Fatal(http.ListenAndServe(":8080", r))
}
