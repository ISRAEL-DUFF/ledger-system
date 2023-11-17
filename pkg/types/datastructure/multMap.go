package datastructure

import "sync"

type MapItem[T any] struct {
	item         T
	multiplicity int32
}

type MultiMap[T comparable, V any] struct {
	mu sync.RWMutex

	keyMap map[T]bool
	data   map[T]MapItem[V]
}

func NewMultiMap[T comparable, V any]() *MultiMap[T, V] {
	return &MultiMap[T, V]{
		data: make(map[T]MapItem[V]),
		mu:   sync.RWMutex{},
	}
}

func (set *MultiMap[T, V]) Set(key T, value V) {
	if set.KeyExists(key) {
		set.mu.Lock()
		v := set.data[key]
		v.multiplicity += 1
		set.data[key] = v
		set.mu.Unlock()
	} else {
		set.mu.Lock()
		set.data[key] = MapItem[V]{
			item:         value,
			multiplicity: 1,
		}
		set.mu.Unlock()
	}
}

func (set *MultiMap[T, V]) Get(key T) (V, bool) {
	set.mu.Lock()
	defer set.mu.Unlock()

	d, exists := set.data[key]
	if exists {
		return d.item, true
	}

	return MapItem[V]{}.item, false
}

func (set *MultiMap[T, V]) KeyExists(key T) bool {
	set.mu.Lock()
	defer set.mu.Unlock()

	_, exists := set.keyMap[key]

	return exists
}

func (set *MultiMap[T, V]) Update(key T, value V) {
	set.mu.Lock()
	defer set.mu.Unlock()

	d, exists := set.data[key]
	if exists {
		d.item = value
		set.data[key] = d
	}
}

func (set *MultiMap[T, V]) Add(key T) {
	set.mu.Lock()
	defer set.mu.Unlock()

	d, exists := set.data[key]
	if exists {
		d.multiplicity += 1
		set.data[key] = d
	}
}

func (set *MultiMap[T, V]) Remove(key T) (V, bool) {
	set.mu.Lock()
	defer set.mu.Unlock()

	d, exists := set.data[key]
	if exists {
		d.multiplicity -= 1

		if d.multiplicity == 0 {
			delete(set.data, key)
		}

		return d.item, true
	}

	return MapItem[V]{}.item, false
}

func (set *MultiMap[T, V]) UpdateAndIncrease(key T, value V) {
	set.updateAndMultiply(key, value, 1)
}

func (set *MultiMap[T, V]) UpdateAndDecrease(key T, value V) {
	set.updateAndMultiply(key, value, -1)
}

func (set *MultiMap[T, V]) updateAndMultiply(key T, item V, frequency int32) {
	set.mu.Lock()
	defer set.mu.Unlock()

	d, exists := set.data[key]

	if exists {
		d.multiplicity += frequency
		d.item = item
		set.data[key] = d
	}
}
