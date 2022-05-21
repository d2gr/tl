package tl

func Contains[T comparable](vs []T, e T) bool {
	for i := range vs {
		if vs[i] == e {
			return true
		}
	}

	return false
}

func SearchFn[T any](vs []T, cmpFn CompareFunc[T]) int {
	for i := range vs {
		if cmpFn(vs[i]) {
			return i
		}
	}

	return -1
}

func ContainsFn[T any](vs []T, cmpFn CompareFunc[T]) bool {
	for i := range vs {
		if cmpFn(vs[i]) {
			return true
		}
	}

	return false
}

func Map[T, E any](set []T, fn func(T) E) []E {
	r := make([]E, len(set))
	for i := range set {
		r[i] = fn(set[i])
	}

	return r
}

func Filter[T any](set []T, cmpFn CompareFunc[T]) []T {
	r := make([]T, 0)
	for i := range set {
		if cmpFn(set[i]) {
			r = append(r, set[i])
		}
	}

	return r
}

func FilterInPlace[T any](set []T, cmpFn CompareFunc[T]) []T {
	for i := 0; i < len(set); i++ {
		if !cmpFn(set[i]) {
			set = append(set[:i], set[i+1:]...)
			i--
		}
	}

	return set
}

func Delete[T comparable](set []T, value T) []T {
	for i := 0; i < len(set); i++ {
		if set[i] == value {
			set = append(set[:i], set[i:]...)
			break
		}
	}

	return set
}
