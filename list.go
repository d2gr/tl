package tl

type listElement[T any] struct {
	value T
	prev  *listElement[T]
	next  *listElement[T]
}

type List[T any] struct {
	prev *listElement[T]
	next *listElement[T]
	size int
}

func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) PushBack(v T) {
	e := &listElement[T]{
		value: v,
		prev:  list.prev,
	}

	if list.prev != nil {
		list.prev.next = e
	} else {
		list.next = e
	}

	list.prev = e
	list.size++
}

func (list *List[T]) PushFront(v T) {
	e := &listElement[T]{
		value: v,
		next:  list.next,
	}

	if list.next != nil {
		list.next.prev = e
	} else {
		list.prev = e
	}

	list.next = e
	list.size++
}

func (list *List[T]) PopFront() (opt OptionalPtr[T]) {
	if list.next != nil {
		opt.Set(&list.next.value)

		list.next = list.next.next
		if list.next != nil {
			list.next.prev = nil
		}

		list.size--
		if list.size == 0 {
			list.prev = nil
		}
	}

	return
}

func (list *List[T]) PopBack() (opt OptionalPtr[T]) {
	if list.prev != nil {
		opt.Set(&list.prev.value)

		list.prev = list.prev.prev
		if list.prev != nil {
			list.prev.next = nil
		}

		list.size--
		if list.size == 0 {
			list.next = nil
		}
	}

	return
}

func (list *List[T]) Reset() {
	list.next = nil
	list.prev = nil
	list.size = 0
}

type forwardIterList[T any] struct {
	next    *listElement[T]
	current *listElement[T]
}

func (list *List[T]) ForwardIter() Iter[T] {
	return &forwardIterList[T]{
		next: list.next,
	}
}

func (iter *forwardIterList[T]) Next() bool {
	iter.current = iter.next
	if iter.next != nil {
		iter.next = iter.next.next
	}

	return iter.current != nil
}

func (iter *forwardIterList[T]) Get() T {
	return iter.current.value
}

func (iter *forwardIterList[T]) GetPtr() *T {
	return &iter.current.value
}

type reverseIterList[T any] struct {
	prev    *listElement[T]
	current *listElement[T]
}

func (list *List[T]) ReverseIter() Iter[T] {
	return &reverseIterList[T]{
		prev: list.prev,
	}
}

func (iter *reverseIterList[T]) Next() bool {
	iter.current = iter.prev
	if iter.prev != nil {
		iter.prev = iter.prev.prev
	}

	return iter.current != nil
}

func (iter *reverseIterList[T]) Get() T {
	return iter.current.value
}

func (iter *reverseIterList[T]) GetPtr() *T {
	return &iter.current.value
}
