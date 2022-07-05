package gots_test

import (
	"testing"

	"github.com/m-mizutani/gots/rands"
	"github.com/stretchr/testify/assert"
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
