package iter

type iterNth[T any] struct {
	inner Iter[T]
	val   *T
}

func (iter *iterNth[T]) Next() bool {
	if iter.inner == nil || !iter.inner.Next() {
		return false
	}

	iter.val = iter.inner.GetPtr()
	iter.inner = nil

	return true
}

func (iter *iterNth[T]) Get() T {
	return *iter.val
}

func (iter *iterNth[T]) GetPtr() *T {
	return iter.val
}

func Nth[T any](inner Iter[T], nth int) Iter[T] {
	for nth != 0 {
		nth--
		inner.Next()
	}

	return &iterNth[T]{
		inner: inner,
	}
}
