package gots_test

import (
	"testing"

	"github.com/m-mizutani/gots"
	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	assert.Len(t, gots.RandomString(6), 6)
	assert.Len(t, gots.RandomString(12), 12)

	// this test has chance 1e-16% of failure
	s := gots.RandomString(128, "晴耕雨読")
	assert.Contains(t, s, "晴")
	assert.Contains(t, s, "耕")
	assert.Contains(t, s, "雨")
	assert.Contains(t, s, "読")
}
