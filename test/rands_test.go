package gots_test

import (
	"testing"

	"github.com/m-mizutani/gots/rands"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandString(t *testing.T) {
	assert.Len(t, rands.New(6), 6)
	assert.Len(t, rands.New(12), 12)

	// this test has chance 1e-16% of failure
	s := rands.New(128, "晴耕雨読")
	assert.Contains(t, s, "晴")
	assert.Contains(t, s, "耕")
	assert.Contains(t, s, "雨")
	assert.Contains(t, s, "読")
}

func TestRandomUint64(t *testing.T) {
	numbers, err := rands.RandomUint64Array(16)
	require.NoError(t, err)
	assert.Len(t, numbers, 16)
	for i := 1; i < 16; i++ {
		assert.NotEqual(t, numbers[0], numbers[i])
	}
}
