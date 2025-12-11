package storage

import "CLI_TodoApp/internal/todo"

type Storage interface {
	Load() ([]todo.Item, error)
	Save([]todo.Item) error
}
