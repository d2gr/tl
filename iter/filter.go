package iter

import "github.com/d2gr/tl"

type iterFilter[T any] struct {
	inner Iter[T]
	fn    tl.CompareFunc[T]
}

func (iter *iterFilter[T]) Next() bool {
	for iter.inner.Next() {
		if iter.fn(iter.inner.Get()) {
			return true
		}
	}

	return false
}

func (iter *iterFilter[T]) Get() T {
	return iter.inner.Get()
}

func (iter *iterFilter[T]) GetPtr() *T {
	return iter.inner.GetPtr()
}

func Filter[T any](inner Iter[T], fn tl.CompareFunc[T]) Iter[T] {
	return &iterFilter[T]{
		inner: inner,
		fn:    fn,
	}
}
