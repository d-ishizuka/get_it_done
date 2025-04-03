package repository

import (
    "gorm.io/gorm"
    "github.com/d-ishizuka/get_it_done/internal/model"
)

type TodoRepository struct {
    db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
    return &TodoRepository{db: db}
}

func (r *TodoRepository) CreateTodo(todo *model.Todo) error {
    return r.db.Create(todo).Error
}

func (r *TodoRepository) GetAllTodos() ([]model.Todo, error) {
    var todos []model.Todo
    err := r.db.Find(&todos).Error
    return todos, err
}

func (r *TodoRepository) GetByID(id uint) (*model.Todo, error) {
    var todo model.Todo
    err := r.db.First(&todo, id).Error
    return &todo, err
}

func (r *TodoRepository) UpdateTodo(todo *model.Todo) error {
    return r.db.Save(todo).Error
}

func (r *TodoRepository) DeleteTodo(id int) error {
    return r.db.Delete(&model.Todo{}, id).Error
}