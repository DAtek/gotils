package gotils

import (
	"time"
)

const ErrorTimeout Error = "TIMEOUT_ERROR"

func NewTimeout(duration time.Duration) (cancel func(), errorChannel <-chan error) {
	errorCh := make(chan error, 1)
	timer := time.AfterFunc(duration, func() { errorCh <- ErrorTimeout })
	cancel = func() { timer.Stop() }
	return cancel, errorCh
}
