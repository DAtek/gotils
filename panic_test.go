package gotils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilOrPanic(t *testing.T) {
	assert.Panics(t, func() {
		NilOrPanic(errors.New("UNEXPECTED_ERROR"))
	})
}
