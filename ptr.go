package gots

// Ptr returns pointer of t.
func Ptr[T any](t T) *T { return &t }
