package gotils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeout(t *testing.T) {
	t.Run("Timeout returns error after specified time elapsed", func(t *testing.T) {
		timeout := NewTimeoutMs(1)

		err := <-timeout.ErrorCh

		assert.EqualError(t, err, ErrorTimeout.Error())
	})

	t.Run("No error if timeout is cancelled", func(t *testing.T) {
		timeout := NewTimeoutMs(1)
		var err error
		go func() { err = <-timeout.ErrorCh }()

		time.Sleep(50 * time.Nanosecond)
		timeout.Cancel()
		time.Sleep(1 * time.Millisecond)

		assert.Nil(t, err)
	})
}
