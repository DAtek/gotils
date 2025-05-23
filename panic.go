package gotils

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}

	return res
}
