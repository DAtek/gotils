package gotils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoroGroup(t *testing.T) {
	t.Run("Group runs all goroutines and can be used multiple times", func(t *testing.T) {
		list := []string{}
		ch := make(chan string)

		publish := func() {
			ch <- "a"
		}

		consume := func() {
			list = append(list, <-ch)
			list = append(list, <-ch)
		}

		group := NewGoroGroup()

		group.Add(publish)
		group.Add(publish)
		group.Add(consume)
		group.Run()

		assert.Equal(t, []string{"a", "a"}, list)

		list = []string{}

		group.Add(publish)
		group.Add(publish)
		group.Add(consume)
		group.Run()

		assert.Equal(t, []string{"a", "a"}, list)
	})
}
