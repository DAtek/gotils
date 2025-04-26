package gotils

import (
	"iter"
)

func Reduce[T any](arr []T, f reduceFunc[T]) T {
	accumulator := arr[0]

	for _, item := range arr[1:] {
		accumulator = f(accumulator, item)
	}

	return accumulator
}

func ReduceSeq[T any](seq iter.Seq[T], f reduceFunc[T]) T {
	var accumulator T
	next, _ := iter.Pull(seq)

	nextVal, valid := next()
	if !valid {
		return accumulator
	}

	accumulator = nextVal

	for {
		nextVal, valid = next()
		if !valid {
			break
		}
		accumulator = f(accumulator, nextVal)
	}

	return accumulator
}

type reduceFunc[T any] func(accumulator T, v T) T
