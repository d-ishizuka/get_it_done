package validator

import (
	"errors"
	"regexp"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func ValidateTodo(todo Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}

	if len(todo.Title) > 100 {
		return errors.New("title must be less than 100 characters")
	}

	if !isValidDescription(todo.Description) {
		return errors.New("description contains invalid characters")
	}

	return nil
}

func isValidDescription(description string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9\s.,!?]*$`)
	return re.MatchString(description)
}