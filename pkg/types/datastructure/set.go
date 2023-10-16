package datastructure

type Set[T comparable] struct {
	keyMap map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		keyMap: make(map[T]bool),
	}
}

func (set *Set[T]) Add(key T) {
	set.keyMap[key] = true
}

func (set *Set[T]) Delete(key T) {
	delete(set.keyMap, key)
}

func (set *Set[T]) AddMany(other Set[T]) {
	for d := range other.keyMap {
		set.keyMap[d] = true
	}
}

func (set *Set[T]) Exists(key T) bool {
	_, exists := set.keyMap[key]

	return exists
}

func (set *Set[T]) Union(set1, set2 Set[T]) *Set[T] {
	newSet := NewSet[T]()

	for key := range set1.keyMap {
		newSet.keyMap[key] = true
	}

	for key := range set2.keyMap {
		newSet.keyMap[key] = true
	}

	return newSet
}

func (set *Set[T]) Intersection(set1, set2 Set[T]) *Set[T] {
	newSet := NewSet[T]()

	for key := range set1.keyMap {
		if set2.keyMap[key] {
			newSet.keyMap[key] = true
		}
	}

	return newSet
}

func (set *Set[T]) Values() []T {
	values := make([]T, 0)

	for i := range set.keyMap {
		values = append(values, i)
	}

	return values
}
