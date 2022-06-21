package gots_test

import (
	"fmt"

	"github.com/m-mizutani/gots"
)

func hello() {
	fmt.Println(gots.GetCaller())
}

func ExampleGetCaller() {
	hello()
	// Output: github.com/m-mizutani/gots/test_test.ExampleGetCaller
}
