package tl

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a, b T) (r T) {
	if a < b {
		r = b
	} else {
		r = a
	}

	return
}

func Min[T constraints.Ordered](a, b T) (r T) {
	if a < b {
		r = a
	} else {
		r = b
	}

	return
}
