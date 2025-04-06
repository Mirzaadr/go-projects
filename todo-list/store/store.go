package store

type Store[T any] struct {
	Filename string
}

func NewStore[T any](filename string) *Store[T] {
	return &Store[T]{Filename: filename}
}

func (s *Store[T]) Save(data T) {}
func (s *Store[T]) Load(data T) {}
