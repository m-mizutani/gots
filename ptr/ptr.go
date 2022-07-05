package ptr

// From returns value of t. If t is nil, it returns empty value of T
func From[T any](t *T) T {
	if t == nil {
		v := new(T)
		return *v
	}
	return *t
}

// To returns pointer of t.
func To[T any](t T) *T { return &t }
