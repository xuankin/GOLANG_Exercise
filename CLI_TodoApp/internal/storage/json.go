package storage

import (
	"CLI_TodoApp/internal/todo"
	"encoding/json"
	"os"
)

type JSONStore struct {
	FilePath string
}

func NewJSONStore(filePath string) *JSONStore {
	return &JSONStore{
		FilePath: filePath,
	}
}
func (s *JSONStore) Load() ([]todo.Item, error) {
	data, err := os.ReadFile(s.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []todo.Item{}, nil
		}
		return nil, err
	}
	var todos []todo.Item
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
func (s *JSONStore) Save(todos []todo.Item) error {
	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FilePath, data, 0644)
}
