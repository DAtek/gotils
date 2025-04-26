package gotils_test

import (
	"iter"
	"testing"

	"github.com/DAtek/gotils"
	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	// given
	list := []int{1, 2, 3}

	// when
	result := gotils.Reduce(list, add)

	// then
	assert.Equal(t, 6, result)
	assert.Equal(t, []int{1, 2, 3}, list)
}

func TestReduceSeq(t *testing.T) {
	t.Run("Simple case", func(t *testing.T) {
		// given
		var generator iter.Seq[int] = func(yield func(int) bool) {
			for i := range 4 {
				if !yield(i) {
					return
				}
			}
		}

		// when
		result := gotils.ReduceSeq(generator, add)

		// then
		assert.Equal(t, 6, result)
	})

	t.Run("Generator doesn't return anything", func(t *testing.T) {
		// given
		var generator iter.Seq[int] = func(yield func(int) bool) {}

		// when
		result := gotils.ReduceSeq(generator, add)

		// then
		assert.Equal(t, 0, result)
	})
}

func add(a, b int) int {
	return a + b
}
