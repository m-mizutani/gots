package examples_test

import (
	"fmt"

	"github.com/m-mizutani/gots"
)

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
