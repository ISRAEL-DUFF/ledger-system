package datastructure

type MapItem[T any] struct {
	item         T
	multiplicity int32
}

type MultiMap[T comparable, V any] struct {
	keyMap map[T]bool

	data map[T]MapItem[V]
}

func NewMap[T comparable, V any]() *MultiMap[T, V] {
	return &MultiMap[T, V]{
		data: make(map[T]MapItem[V]),
	}
}

func (set *MultiMap[T, V]) Set(key T, value V) {
	if set.KeyExists(key) {
		v := set.data[key]
		v.multiplicity += 1
		set.data[key] = v
	} else {
		set.data[key] = MapItem[V]{
			item:         value,
			multiplicity: 1,
		}
	}
}

func (set *MultiMap[T, V]) Get(key T) (V, bool) {
	d, exists := set.data[key]
	if exists {
		return d.item, true
	}

	return MapItem[V]{}.item, false
}

func (set *MultiMap[T, V]) KeyExists(key T) bool {
	_, exists := set.keyMap[key]

	return exists
}

func (set *MultiMap[T, V]) Update(key T, value V) {
	d, exists := set.data[key]
	if exists {
		d.item = value
		set.data[key] = d
	}
}

func (set *MultiMap[T, V]) Add(key T) {
	d, exists := set.data[key]
	if exists {
		d.multiplicity += 1
		set.data[key] = d
	}
}

func (set *MultiMap[T, V]) Remove(key T) (V, bool) {
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
	d, exists := set.data[key]

	if exists {
		d.multiplicity += frequency
		d.item = item
		set.data[key] = d
	}
}
