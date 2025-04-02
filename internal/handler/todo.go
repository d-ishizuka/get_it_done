package handler

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/d-ishizuka/get_it_done/internal/model"
)

type TodoRepo interface {
    CreateTodo(todo *model.Todo) error
    GetAllTodos() ([]model.Todo, error)
    UpdateTodo(todo *model.Todo) error
    DeleteTodo(id int) error
}

var repo TodoRepo

func Initialize(todoRepo TodoRepo) {
    repo = todoRepo
}

// CreateTodo handles creation of a new todo item
func CreateTodo(w http.ResponseWriter, r *http.Request) {
    var todo model.Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := repo.CreateTodo(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(todo)
}

// GetTodos returns all todo items
func GetTodos(w http.ResponseWriter, r *http.Request) {
    todos, err := repo.GetAllTodos()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

// UpdateTodo updates an existing todo item
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid todo ID", http.StatusBadRequest)
        return
    }

    var todo model.Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    todo.ID = id
    if err := repo.UpdateTodo(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

// DeleteTodo deletes a todo item
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid todo ID", http.StatusBadRequest)
        return
    }

    if err := repo.DeleteTodo(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}