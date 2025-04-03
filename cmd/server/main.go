package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "github.com/d-ishizuka/get_it_done/internal/handler"
    "github.com/d-ishizuka/get_it_done/internal/model"
    "github.com/d-ishizuka/get_it_done/internal/repository"
)

func main() {
    // データベース接続
    db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // テーブル自動作成
    db.AutoMigrate(&model.Todo{})
    
    // リポジトリとハンドラーの初期化
    todoRepo := repository.NewTodoRepository(db)
    handler.Initialize(todoRepo)

    r := mux.NewRouter()

    // API routes
    r.HandleFunc("/todos", handler.CreateTodo).Methods("POST")
    r.HandleFunc("/todos", handler.GetTodos).Methods("GET")
    r.HandleFunc("/todos/{id}", handler.UpdateTodo).Methods("PUT")
    r.HandleFunc("/todos/{id}", handler.CreateTodo).Methods("DELETE")

    // Start the server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}