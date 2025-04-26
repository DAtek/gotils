package gotils

import (
	"time"
)

func MustDateTimeFromString(dateTimeStr string) *time.Time {
	return Must(DateTimeFromString(dateTimeStr))
}

func DateTimeFromString(dateTimeStr string) (*time.Time, error) {
	time, err := time.Parse(time.DateTime, dateTimeStr)
	if err != nil {
		return nil, err
	}
	return &time, nil
}
