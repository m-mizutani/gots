package gots_test

import (
	"fmt"

	"github.com/m-mizutani/gots/stack"
)

func hello() {
	fmt.Println(stack.GetCaller())
}

func ExampleGetCaller() {
	hello()
	// Output: github.com/m-mizutani/gots/test_test.ExampleGetCaller
}
