package gotils

import (
	"time"
)

const ErrorTimeout Error = "TIMEOUT_ERROR"

type timeout struct {
	ErrorCh  chan error
	timer    *time.Timer
	duration time.Duration
}

func NewTimeout(duration time.Duration) *timeout {
	errorCh := make(chan error, 1)
	t := &timeout{
		ErrorCh:  errorCh,
		duration: duration,
		timer: time.AfterFunc(duration, func() {
			errorCh <- ErrorTimeout
		}),
	}

	return t
}

func (t *timeout) Cancel() {
	t.timer.Stop()
}
