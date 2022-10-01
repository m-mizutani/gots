package set

type Set[T comparable] struct {
	FrozenSet[T]
}

type FrozenSet[T comparable] struct {
	data map[T]struct{}
}

// New creates set
func New[T comparable](init ...T) *Set[T] {
	s := &Set[T]{
		FrozenSet: FrozenSet[T]{
			data: make(map[T]struct{}),
		},
	}

	for _, i := range init {
		s.data[i] = struct{}{}
	}

	return s
}

// Add adds the item to set
func (x *Set[T]) Add(items ...T) {
	for _, item := range items {
		x.data[item] = struct{}{}
	}
}

// Remove deletes the item from set. It returns true if the item exists.
func (x *Set[T]) Remove(item T) bool {
	if _, ok := x.data[item]; ok {
		delete(x.data, item)
		return true
	}

	return false
}

// Clear discards all items from the set
func (x *Set[T]) Clear() {
	x.data = make(map[T]struct{})
}

// Has returns true if the item exists
func (x *FrozenSet[T]) Has(item T) bool {
	_, ok := x.data[item]
	return ok
}

// Len returns size of set
func (x *FrozenSet[T]) Len() int {
	return len(x.data)
}

// Items returns array of items in the set
func (x *FrozenSet[T]) Items() []T {
	a := make([]T, len(x.data))
	var idx int
	for item := range x.data {
		a[idx] = item
		idx++
	}
	return a
}

// Union creates a new set has all items from x and sets
func (x *FrozenSet[T]) Union(sets ...*Set[T]) *Set[T] {
	z := New[T]()

	for v := range x.data {
		z.Add(v)
	}

	for _, s := range sets {
		for v := range s.data {
			z.Add(v)
		}
	}

	return z
}

// Diff creates a new set has items in the set && not including others
func (x *FrozenSet[T]) Diff(others ...*Set[T]) *Set[T] {
	z := x.Copy()

	for _, o := range others {
		for v := range o.data {
			z.Remove(v)
		}
	}

	return z
}

// Intersection creates a new set has items in both of the set and others
func (x *FrozenSet[T]) Intersection(others ...*Set[T]) *Set[T] {
	z := New[T]()

	hasAll := func(v T) bool {
		for _, o := range others {
			if !o.Has(v) {
				return false
			}
		}
		return true
	}

	for i := range x.data {
		if hasAll(i) {
			z.Add(i)
		}
	}

	return z
}

// Freeze provides a new FrozenSet created from original set. Returned FrozenSet is not updated by modifying original set.
func (x *Set[T]) Freeze() *FrozenSet[T] {
	f := &FrozenSet[T]{
		data: make(map[T]struct{}),
	}
	for d := range x.data {
		f.data[d] = struct{}{}
	}
	return f
}

// Copy provides a new shallow copied set.
func (x *FrozenSet[T]) Copy() *Set[T] {
	s := &Set[T]{
		FrozenSet: FrozenSet[T]{
			data: make(map[T]struct{}),
		},
	}
	for d := range x.data {
		s.data[d] = struct{}{}
	}
	return s
}
