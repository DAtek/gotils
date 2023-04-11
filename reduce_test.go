package gotils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	list := []int{1, 2, 3}

	add := func(a, b int) int {
		return a + b
	}

	result := Reduce(list, add)

	assert.Equal(t, 6, result)
	assert.Equal(t, []int{1, 2, 3}, list)
}
