package gotils

// Useful in tests
func NilOrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// Useful in tests
func ResultOrPanic[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}

	return res
}
