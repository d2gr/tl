package tl

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](numbers ...T) (r T) {
	if len(numbers) == 0 {
		return
	}

	r = numbers[0]

	for _, v := range numbers[1:] {
		if v > r {
			r = v
		}
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
