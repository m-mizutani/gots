package gots

func Map[T1, T2 any](input []T1, f func(T1) T2) []T2 {
	output := make([]T2, len(input))
	for i := range input {
		output[i] = f(input[i])
	}
	return output
}

func Filter[T any](input []T, f func(T) bool) []T {
	var output []T
	for i := range input {
		if f(input[i]) {
			output = append(output, input[i])
		}
	}

	return output
}

func Reduce[T any](input []T, f func(prev T, current T) T, init T) T {
	c := init
	for i := range input {
		c = f(input[i], c)
	}
	return c
}

func Unique[T comparable](input []T) []T {
	var output []T
	added := map[T]struct{}{}
	for i := range input {
		if _, ok := added[input[i]]; ok {
			continue
		}
		output = append(output, input[i])
		added[input[i]] = struct{}{}
	}

	return output
}

func Count[T comparable](input []T, target T) int {
	var c int
	for i := range input {
		if target == input[i] {
			c++
		}
	}
	return c
}

func CountIf[T any](input []T, f func(v T) bool) int {
	var c int
	for i := range input {
		if f(input[i]) {
			c++
		}
	}
	return c
}
