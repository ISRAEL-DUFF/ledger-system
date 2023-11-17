package datastructure

import "sync"

type Map[T comparable, V any] struct {
	mu   sync.RWMutex
	data map[T]V
}

func NewMap[T comparable, V any]() *Map[T, V] {
	return &Map[T, V]{
		data: make(map[T]V),
		mu:   sync.RWMutex{},
	}
}

func (set *Map[T, V]) Set(key T, value V) {
	set.mu.Lock()
	set.data[key] = value
	set.mu.Unlock()
}

func (set *Map[T, V]) Get(key T) (V, bool) {
	set.mu.Lock()
	defer set.mu.Unlock()

	d, exists := set.data[key]

	return d, exists
}

// func (set *Map[T, V]) KeyExists(key T) bool {
// 	set.mu.Lock()
// 	defer set.mu.Unlock()

// 	_, exists := set.keyMap[key]

// 	return exists
// }

func (set *Map[T, V]) Remove(key T) bool {
	set.mu.Lock()
	defer set.mu.Unlock()

	delete(set.data, key)

	return true
}
