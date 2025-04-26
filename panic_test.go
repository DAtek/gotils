package gotils_test

import (
	"errors"
	"testing"

	"github.com/DAtek/gotils"
	"github.com/stretchr/testify/assert"
)

func TestNoErr(t *testing.T) {
	assert.Panics(t, func() {
		gotils.NoErr(errors.New("UNEXPECTED_ERROR"))
	})
}
