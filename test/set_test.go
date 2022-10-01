package gots_test

import (
	"testing"

	"github.com/m-mizutani/gots/set"
	"github.com/m-mizutani/gots/slice"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		s := set.New("a", "b")

		assert.True(t, s.Has("a"))
		assert.True(t, s.Has("b"))
		assert.False(t, s.Has("c"))
		assert.Equal(t, 2, s.Len())

		s.Add("c", "d")

		assert.True(t, s.Has("a"))
		assert.True(t, s.Has("b"))
		assert.True(t, s.Has("c"))
		assert.True(t, s.Has("d"))
		assert.Equal(t, 4, s.Len())

		assert.True(t, s.Remove("a"))
		assert.False(t, s.Remove("x"))

		assert.Equal(t, 3, s.Len())
		assert.False(t, s.Has("a"))

		items := s.Items()
		assert.Len(t, items, 3)
		assert.False(t, slice.Contains(items, "a"))
		assert.True(t, slice.Contains(items, "b"))
		assert.True(t, slice.Contains(items, "c"))
		assert.True(t, slice.Contains(items, "d"))

		s.Clear()
		assert.Equal(t, 0, s.Len())
		assert.False(t, s.Has("b"))
		assert.False(t, s.Has("c"))
		assert.False(t, s.Has("d"))
	})

	t.Run("copied set is not affected by changes in original set", func(t *testing.T) {
		s1 := set.New(1, 2, 5)
		s2 := s1.Copy()
		assert.True(t, s1.Remove(1))

		assert.False(t, s1.Has(1))
		assert.Equal(t, 2, s1.Len())
		assert.True(t, s2.Has(1))
		assert.Equal(t, 3, s2.Len())
	})

	t.Run("frozen set is not affected by changes in original set", func(t *testing.T) {
		s1 := set.New(1, 2, 5)
		s2 := s1.Freeze()
		assert.True(t, s1.Remove(1))

		assert.False(t, s1.Has(1))
		assert.Equal(t, 2, s1.Len())
		assert.True(t, s2.Has(1))
		assert.Equal(t, 3, s2.Len())
	})

	t.Run("Intersection returns another set", func(t *testing.T) {
		s1 := set.New(1, 2, 3, 4, 5)
		s2 := s1.Intersection(set.New(3, 4, 5), set.New(3, 4))
		assert.Equal(t, 2, s2.Len())
		assert.False(t, s2.Has(1))
		assert.False(t, s2.Has(2))
		assert.True(t, s2.Has(3))
		assert.True(t, s2.Has(4))
		assert.False(t, s2.Has(5))

		assert.True(t, s1.Remove(3))
		assert.True(t, s2.Has(3)) // not affected
	})

	t.Run("Diff returns another set", func(t *testing.T) {
		s1 := set.New(1, 2, 3, 4)
		s2 := s1.Diff(set.New(3, 4, 5))
		assert.Equal(t, 2, s2.Len())
		assert.True(t, s2.Has(1))
		assert.True(t, s2.Has(2))
		assert.False(t, s2.Has(3))
		assert.False(t, s2.Has(4))
		assert.False(t, s2.Has(5))

		assert.True(t, s1.Remove(2))
		assert.True(t, s2.Has(2)) // not affected
	})

	t.Run("Union returns another set", func(t *testing.T) {
		s1 := set.New(1, 2, 3, 4)
		s2 := s1.Union(set.New(3, 4, 5))
		assert.Equal(t, 5, s2.Len())
		assert.True(t, s2.Has(1))
		assert.True(t, s2.Has(2))
		assert.True(t, s2.Has(3))
		assert.True(t, s2.Has(4))
		assert.True(t, s2.Has(5))

		assert.True(t, s1.Remove(2))
		assert.True(t, s2.Has(2)) // not affected
	})
}
