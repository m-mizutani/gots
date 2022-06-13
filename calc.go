package gots

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[T Number](input []T) T {
	sum := new(T)
	for i := range input {
		*sum = input[i] + *sum
	}
	return *sum
}

func Avg[T Number](input []T) T {
	sum := new(T)
	cnt := new(T)
	for i := range input {
		*cnt = *cnt + 1
		*sum = input[i] + *sum
	}
	return *sum / *cnt
}
