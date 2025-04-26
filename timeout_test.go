package gotils_test

import (
	"testing"
	"time"

	"github.com/DAtek/gotils"
	"github.com/stretchr/testify/assert"
)

func TestTimeout(t *testing.T) {
	t.Run("Timeout returns error after specified time elapsed", func(t *testing.T) {
		_, errorChannel := gotils.NewTimeout(1 * time.Millisecond)

		err := <-errorChannel

		assert.EqualError(t, err, gotils.ErrorTimeout.Error())
	})

	t.Run("No error if timeout is cancelled", func(t *testing.T) {
		cancel, errorChannel := gotils.NewTimeout(1 * time.Millisecond)
		var err error
		go func() { err = <-errorChannel }()

		time.Sleep(50 * time.Nanosecond)
		cancel()
		time.Sleep(1 * time.Millisecond)

		assert.Nil(t, err)
	})
}
