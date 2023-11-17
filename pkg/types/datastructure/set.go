package datastructure

import "sync"

type Set[T comparable] struct {
	mu        sync.RWMutex
	keyMap    map[T]bool
	itemCount int
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		keyMap:    make(map[T]bool),
		itemCount: 0,
		mu:        sync.RWMutex{},
	}
}

func (set *Set[T]) Add(key T) {
	if !set.Exists(key) {
		set.itemCount += 1
	}

	set.mu.Lock()
	defer set.mu.Unlock()

	set.keyMap[key] = true
}

func (set *Set[T]) Delete(key T) {
	if set.Exists(key) {
		set.itemCount -= 1
	}

	set.mu.Lock()
	defer set.mu.Unlock()

	delete(set.keyMap, key)
}

func (set *Set[T]) AddMany(other *Set[T]) {
	for d := range other.keyMap {
		if !set.Exists(d) {
			set.mu.Lock()
			set.itemCount += 1
			set.mu.Unlock()
		}

		set.mu.Lock()
		set.keyMap[d] = true
		set.mu.Unlock()
	}
}

func (set *Set[T]) Exists(key T) bool {
	set.mu.RLock()
	defer set.mu.RUnlock()

	_, exists := set.keyMap[key]

	return exists
}

func (set *Set[T]) Union(otherSet *Set[T]) *Set[T] {
	set.mu.Lock()
	defer set.mu.Unlock()

	newSet := NewSet[T]()

	for key := range otherSet.keyMap {
		newSet.Add(key)
	}

	return newSet
}

func (set *Set[T]) Intersection(otherSet *Set[T]) *Set[T] {
	set.mu.Lock()
	defer set.mu.Unlock()

	newSet := NewSet[T]()

	for key := range set.keyMap {
		if otherSet.Exists(key) {
			newSet.Add(key)
		}
	}

	return newSet
}

func (set *Set[T]) Difference(otherSet *Set[T]) *Set[T] {
	set.mu.Lock()
	defer set.mu.Unlock()

	newSet := NewSet[T]()

	for key := range set.keyMap {
		if !otherSet.Exists(key) {
			newSet.Add(key)
		}
	}

	return newSet
}

func (set *Set[T]) Values() []T {
	set.mu.RLock()
	defer set.mu.RUnlock()

	values := make([]T, 0)

	for i := range set.keyMap {
		values = append(values, i)
	}

	return values
}

func (set *Set[T]) Len() int {
	return set.itemCount
}
