package transform

type MapClosure[T any] func(T) T
type FilterClosure[T any] func(T) bool

func Map[T any](slice []T, f MapClosure[T]) []T {
	out := make([]T, len(slice))

	for i, item := range slice {
		out[i] = f(item)
	}

	return out
}

func Filter[T any](slice []T, f FilterClosure[T]) []T {
	out := []T{}

	for k, v := range slice {
		if f(v) {
			out[k] = v
		}
	}

	return out
}

func First[T any](slice []T, f FilterClosure[T]) T {
	var out T

	for _, v := range slice {
		if f(v) {
			return v
		}
	}

	return out
}

func Reduce[T, A any](slice []T, f func(A, T) A, accumulator A) A {
	for _, v := range slice {
		accumulator = f(accumulator, v)
	}

	return accumulator
}