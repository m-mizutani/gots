package gots_test

import (
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
