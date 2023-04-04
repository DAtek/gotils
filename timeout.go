package gotils

import (
	"time"
)

const ErrorTimeout = Error("TIMEOUT_ERROR")

type timeout struct {
	finished bool
	duration time.Duration
	ErrorCh  chan error
}

func NewTimeout(duration time.Duration) *timeout {
	t := &timeout{
		duration: duration,
		ErrorCh:  make(chan error, 1),
	}

	t.start()
	return t
}

// Useful in tests
func NewTimeoutMs(durationMs uint) *timeout {
	t := NewTimeout(time.Duration(durationMs) * time.Millisecond)
	t.start()
	return t
}

func (t *timeout) Cancel() {
	t.finished = true
}

func (t *timeout) start() {
	go func() {
		time.Sleep(t.duration)
		if t.finished {
			return
		}
		t.ErrorCh <- ErrorTimeout
	}()
}
