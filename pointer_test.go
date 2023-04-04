package gotils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	text := Pointer("asd")

	assert.Equal(t, "asd", *text)
}
