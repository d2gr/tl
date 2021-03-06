package iter

import (
	"github.com/d2gr/tl"
	"golang.org/x/exp/constraints"
)

type iterRange[T constraints.Integer] struct {
	start, stop, step T
	current           T
	value             T
}

func (iter *iterRange[T]) Next() bool {
	iter.value = iter.start + iter.step*iter.current
	iter.current++

	if iter.step > 0 {
		return iter.value < iter.stop
	}

	return iter.value > iter.stop
}

func (iter *iterRange[T]) Get() T {
	return iter.value
}

func (iter *iterRange[T]) GetPtr() *T {
	return &iter.value
}

func Range[T constraints.Integer](start, stop, step T) tl.Iter[T] {
	return &iterRange[T]{
		start: start, stop: stop, step: step,
		current: 0,
	}
}
