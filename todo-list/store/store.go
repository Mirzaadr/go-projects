package store

import (
	"encoding/json"
	"os"
)

type Store[T any] struct {
	Filename string
}

func NewStore[T any](filename string) *Store[T] {
	return &Store[T]{Filename: filename}
}

func (s *Store[T]) Save(data T) error {
	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(s.Filename, fileData, 0644)
}
func (s *Store[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.Filename)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
