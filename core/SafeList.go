package core

import (
	"errors"
	"sync"
)

type SafeList[T any] struct {
	mu   sync.RWMutex
	data map[string]*T
}

func (s *SafeList[T]) Get(name string) *T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := s.data[name]

	return result
}

func (s *SafeList[T]) Add(key string, obj *T) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	search := s.data[key]
	if search != nil {
		return errors.New(key + " already exists")
	}

	s.data[key] = obj

	return nil
}

func (s *SafeList[T]) Remove(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = nil
}
