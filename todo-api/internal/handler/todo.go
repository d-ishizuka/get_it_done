package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "todo-api/internal/model"
    "todo-api/internal/repository"
)

func CreateTodo(c *gin.Context) {
    var todo model.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := repository.CreateTodo(&todo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create todo"})
        return
    }
    c.JSON(http.StatusCreated, todo)
}

func GetTodos(c *gin.Context) {
    todos, err := repository.GetTodos()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve todos"})
        return
    }
    c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    var todo model.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    todo.ID = id
    if err := repository.UpdateTodo(&todo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update todo"})
        return
    }
    c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    if err := repository.DeleteTodo(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete todo"})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}