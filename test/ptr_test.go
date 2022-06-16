package gots_test

import (
	"time"

	"github.com/m-mizutani/gots"
)

func ExamplePtr() {
	type input struct {
		ts *int64
	}

	_ = input{
		ts: gots.Ptr(time.Now().Unix()),
	}
}
