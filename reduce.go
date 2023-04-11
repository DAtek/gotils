package gotils

func reduce[T any](arr []T, f func(accumulator, currentValue T) T) T {
	accumulator := arr[0]

	for _, item := range arr[1:] {
		accumulator = f(accumulator, item)
	}

	return accumulator
}
