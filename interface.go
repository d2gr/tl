package tl

type Iter[T any] interface {
	Next() bool
	Get() T
	GetPtr() *T
}

func Advance[T any](iter Iter[T], count int) {
	for i := 0; i < count && iter.Next(); i++ {
	}
}
