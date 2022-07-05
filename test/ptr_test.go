package gots_test

import (
	"time"

	"github.com/m-mizutani/gots/ptr"
)

func ExamplePtr() {
	type input struct {
		ts *int64
	}

	_ = input{
		ts: ptr.To(time.Now().Unix()),
	}
}
