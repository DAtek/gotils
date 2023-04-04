package gotils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvConfig(t *testing.T) {
	key := "COLOR"
	value := "green"
	configColor := EnvConfig(key)

	os.Setenv(key, value)

	assert.Equal(t, value, configColor.Load())
}
