package iter

type Iter[T any] interface {
	Next() bool
	Get() T
	GetPtr() *T
}
