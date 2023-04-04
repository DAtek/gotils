package gotils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateTimeFromStringPanic(t *testing.T) {
	t.Run("Returns time if format is valid", func(t *testing.T) {
		dateTimeStr := "2023-01-01 12:00:00"
		dateTime := DateTimeFromStringPanic(dateTimeStr)
		expected, _ := time.Parse(time.DateTime, dateTimeStr)

		assert.Equal(t, expected, *dateTime)
	})

	t.Run("Panics if format is invalid", func(t *testing.T) {
		assert.Panics(t, func() {
			DateTimeFromStringPanic("asd")
		})
	})
}
