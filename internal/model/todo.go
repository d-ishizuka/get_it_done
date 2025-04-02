package model

import "errors"

type Todo struct {
    ID        int   `json:"id" gorm:"primaryKey"`
    Title     string `json:"title" gorm:"not null"`
    Completed bool   `json:"completed" gorm:"default:false"`
}

func (t *Todo) Validate() error {
    if t.Title == "" {
        return errors.New("title is required")
    }
    return nil
}