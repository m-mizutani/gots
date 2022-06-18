package gots_test

import (
	"fmt"
	"testing"

	"github.com/m-mizutani/gots"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	input := []string{"blue", "orange", "red"}
	output := gots.Map(input, func(v string) int {
		return len(v)
	})
	require.Len(t, output, 3)
	assert.Equal(t, 4, output[0])
	assert.Equal(t, 6, output[1])
	assert.Equal(t, 3, output[2])
}

func TestFilter(t *testing.T) {
	out := gots.Filter([]int{1, 2, 3, 4, 5}, func(v int) bool {
		return v%2 == 0
	})
	require.Len(t, out, 2)
	assert.Equal(t, 2, out[0])
	assert.Equal(t, 4, out[1])
}

func TestReduce(t *testing.T) {
	out := gots.Reduce([]int{1, 2, 3, 4, 5}, func(p, c int) int {
		return p + c
	}, 10)
	assert.Equal(t, 25, out)
}

func ExampleMap() {
	out := gots.Map([]string{"blue", "orange", "red"}, func(v string) int {
		return len(v)
	})
	fmt.Println(out)
	// Output: [4 6 3]
}

func ExampleFilter() {
	out := gots.Filter([]int{1, 2, 3, 4, 5}, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(out)
	// Output: [2 4]
}

func ExampleReduce() {
	out := gots.Reduce([]int{1, 2, 3, 4, 5}, func(p, c int) int {
		return p + c
	}, 10)
	fmt.Println(out)
	// Output: 25
}

func ExampleUnique() {
	out := gots.Unique([]string{"A", "B", "A", "C", "B"})
	fmt.Println(out)
	// Output: [A B C]
}

func ExampleCount() {
	out := gots.Count([]string{"A", "B", "A", "C", "B"}, "B")
	fmt.Println(out)
	// Output: 2
}

func ExampleCountIf() {
	out := gots.CountIf([]string{"A", "B", "A", "C", "B"}, func(v string) bool {
		return v == "A" || v == "C"
	})
	fmt.Println(out)
	// Output: 3
}

func ExampleContains() {
	fmt.Println(gots.Contains([]string{"A", "B", "A", "C", "B"}, "B"))
	// Output: true
}
