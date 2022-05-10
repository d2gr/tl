package tl

type listElement[T any] struct {
	value T
	prev  *listElement[T]
	next  *listElement[T]
}

type List[T any] struct {
	root listElement[T]
	size int
}

func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) PushBack(v T) {
	e := &listElement[T]{
		value: v,
		prev:  list.root.prev,
		next:  &list.root,
	}

	if list.root.prev != nil {
		list.root.prev.next = e
	} else {
		list.root.next = e
	}

	if e.prev == nil {
		e.prev = &list.root
	}

	list.root.prev = e
	list.size++
}

func (list *List[T]) PushFront(v T) {
	e := &listElement[T]{
		value: v,
		prev:  &list.root,
		next:  list.root.next,
	}

	if list.root.next != nil {
		list.root.next.prev = e
	} else {
		list.root.prev = e
	}

	if e.next == nil {
		e.next = &list.root
	}

	list.root.next = e
	list.size++
}

func (list *List[T]) PopFront() (opt OptionalPtr[T]) {
	if list.root.next != nil {
		opt.Set(&list.root.next.value)

		list.root.next = list.root.next.next

		list.size--
		if list.size == 0 {
			list.root.prev = nil
		}
	}

	return
}

func (list *List[T]) PopBack() (opt OptionalPtr[T]) {
	if list.root.prev != nil {
		opt.Set(&list.root.prev.value)

		list.root.prev = list.root.prev.prev

		list.size--
		if list.size == 0 {
			list.root.next = nil
		}
	}

	return
}

func (list *List[T]) Reset() {
	list.root.next = nil
	list.root.prev = nil
	list.size = 0
}

type forwardIterList[T any] struct {
	root *listElement[T]
	next *listElement[T]
}

func (list *List[T]) ForwardIter() Iter[T] {
	return &forwardIterList[T]{
		root: &list.root,
		next: &list.root,
	}
}

func (iter *forwardIterList[T]) Next() bool {
	cond := iter.root != iter.next.next && iter.next.next != nil
	iter.next = iter.next.next
	return cond
}

func (iter *forwardIterList[T]) Get() T {
	return iter.next.value
}

func (iter *forwardIterList[T]) GetPtr() *T {
	return &iter.next.value
}

type reverseIterList[T any] struct {
	root *listElement[T]
	next *listElement[T]
}

func (list *List[T]) ReverseIter() Iter[T] {
	return &reverseIterList[T]{
		root: &list.root,
		next: &list.root,
	}
}

func (iter *reverseIterList[T]) Next() bool {
	cond := iter.root != iter.next.prev && iter.next.prev != nil
	iter.next = iter.next.prev
	return cond
}

func (iter *reverseIterList[T]) Get() T {
	return iter.next.value
}

func (iter *reverseIterList[T]) GetPtr() *T {
	return &iter.next.value
}
