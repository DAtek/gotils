package gotils

import (
	"time"
)

// Useful in tests
func DateTimeFromStringPanic(dateTimeStr string) *time.Time {
	return ResultOrPanic(DateTimeFromString(dateTimeStr))
}

func DateTimeFromString(dateTimeStr string) (*time.Time, error) {
	time, err := time.Parse(time.DateTime, dateTimeStr)
	if err != nil {
		return nil, err
	}
	return &time, nil
}
