package iter

type iterWindowCopy[T any] struct {
	size  int
	win   []T
	inner Iter[T]
}

func (iter *iterWindowCopy[T]) Next() bool {
	if len(iter.win) != 0 {
		iter.win = append([]T{}, iter.win[1:]...)
	}

	for len(iter.win) < iter.size {
		if !iter.inner.Next() {
			return false
		}

		iter.win = append(iter.win, iter.inner.Get())
	}

	return true
}

func (iter *iterWindowCopy[T]) Get() []T {
	return iter.win
}

func (iter *iterWindowCopy[T]) GetPtr() *[]T {
	return &iter.win
}

func WindowCopy[T any](inner Iter[T], n int) Iter[[]T] {
	return &iterWindowCopy[T]{
		inner: inner,
		win:   make([]T, 0, n),
		size:  n,
	}
}
