package annotate

import "fmt"


type registry[T any] struct {
	register map[string]T
	getKey func(T) string
}

func NewRegistry[T any](keyFn func(T) string) *registry[T] {
	return &registry[T]{
		register: make(map[string]T),
		getKey: keyFn,
	}
}

func (r *registry[T]) Get(key string) (T, error) {
	if item, exists := r.register[key]; exists {
		return item, nil
	}

	return *new(T), fmt.Errorf("error retrieving item: item with name '%s' does not exist", key)
}

func (r *registry[T]) Register(item T) error {
	key := r.getKey(item)
	if _, err := r.Get(key); err != nil {
		r.register[key] = item
	}

	return fmt.Errorf("error registering item: item with name '%s' already exists", key)
}