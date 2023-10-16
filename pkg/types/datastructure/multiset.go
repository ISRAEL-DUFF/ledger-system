package datastructure

type MultiSet[T comparable] struct {
	keyMap map[T]int32
	count  int32
}

func NewMultiSet[T comparable]() *MultiSet[T] {
	return &MultiSet[T]{
		keyMap: make(map[T]int32),
	}
}

func (set *MultiSet[T]) Add(key T) {
	if set.Exists(key) {
		set.keyMap[key] += 1
	} else {
		set.keyMap[key] = 1
	}

	set.count += 1
}

func (set *MultiSet[T]) Delete(key T) {
	frequency, exists := set.keyMap[key]
	if exists {
		if frequency == 1 {
			delete(set.keyMap, key)
		} else {
			frequency -= 1
			set.keyMap[key] = frequency
		}

		set.count -= 1
	}
}

func (set *MultiSet[T]) Exists(key T) bool {
	_, exists := set.keyMap[key]

	return exists
}

func (set *MultiSet[T]) Union(set1, set2 Set[T]) *MultiSet[T] {
	newSet := NewMultiSet[T]()

	for key := range set1.keyMap {
		newSet.Add(key)
	}

	for key := range set2.keyMap {
		newSet.Add(key)
	}

	return newSet
}

func (set *MultiSet[T]) Intersection(set1, set2 MultiSet[T]) *MultiSet[T] {
	newSet := NewMultiSet[T]()

	for key := range set1.keyMap {
		if set2.Exists(key) {
			newSet.Add(key)

		}
	}

	return newSet
}

func (set *MultiSet[T]) Values() []T {
	values := make([]T, 0)

	for i, freq := range set.keyMap {
		for j := 1; j <= int(freq); j++ {
			values = append(values, i)
		}
	}

	return values
}
