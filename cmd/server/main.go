package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // API routes
    r.HandleFunc("/todos", CreateTodo).Methods("POST")
    r.HandleFunc("/todos", GetTodos).Methods("GET")
    r.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
    r.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

    // Start the server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}