package gots_test

import (
	"testing"

	"github.com/m-mizutani/gots"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 10, gots.Sum([]int{1, 2, 3, 4}))
	assert.Equal(t, float64(12), gots.Sum([]float64{1.5, 2.5, 3.1, 4.9}))
}

func TestAvg(t *testing.T) {
	assert.Equal(t, 2, gots.Avg([]int{1, 2, 3, 4}))
	assert.Equal(t, 3, gots.Avg([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, float64(3), gots.Avg([]float64{1.5, 2.5, 3.1, 4.9}))
}
