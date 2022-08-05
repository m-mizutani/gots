package cache

type Item[T any] struct {
	cached bool
	data   T
}

func NewItem[T any]() *Item[T] {
	return &Item[T]{}
}

func (x *Item[T]) Get() T {
	return x.data
}

func (x *Item[T]) Set(data T) {
	x.data = data
	x.cached = true
}

func (x *Item[T]) HasCache() bool {
	return x.cached
}

func (x *Item[T]) Clear() {
	x.cached = false
	v := new(T)
	x.data = *v
}
