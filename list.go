package tl

import "fmt"

type listElement[T any] struct {
	value T
	prev  *listElement[T]
	next  *listElement[T]
}

func (e *listElement[T]) Drop() {
	if e.prev != nil {
		e.prev.next = e.next
	}

	if e.next != nil {
		e.next.prev = e.prev
	}
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

	if e.prev == nil {
		e.prev = &list.root
	}

	if list.root.prev != nil {
		list.root.prev.next = e
	} else {
		list.root.next = e
	}

	list.root.prev = e
	list.size++
}

func (list *List[T]) PushFront(v T) {
	e := &listElement[T]{
		value: v,
		next:  list.root.next,
		prev:  &list.root,
	}

	if e.next == nil {
		e.next = &list.root
	}

	if list.root.next != nil {
		list.root.next.prev = e
	} else {
		list.root.prev = e
	}

	list.root.next = e
	list.size++
}

func (list *List[T]) Front() (opt OptionalPtr[T]) {
	if list.root.next != nil {
		opt.Set(&list.root.next.value)
	}

	return
}

func (list *List[T]) Back() (opt OptionalPtr[T]) {
	if list.root.prev != nil {
		opt.Set(&list.root.prev.value)
	}

	return
}

func Print[T any](list *listElement[T]) {
	fmt.Printf("%p - %p %p\n", list, list.prev, list.next)
	for next := list.next; next != nil && next != list; next = next.next {
		fmt.Printf("%p (%v) = %p - %p\n", next, next.value, next.prev, next.next)
	}
	println("-------")
}

func (list *List[T]) PopFront() (opt OptionalPtr[T]) {
	if list.root.next != nil {
		opt.Set(&list.root.next.value)

		list.root.next.Drop()

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

		list.root.prev.Drop()
		// list.root.prev = list.root.prev.prev

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
	root    *listElement[T]
	next    *listElement[T]
	current *listElement[T]
}

func (list *List[T]) ForwardIter() IterDrop[T] {
	return &forwardIterList[T]{
		root: &list.root,
		next: list.root.next,
	}
}

func (iter *forwardIterList[T]) Drop() {
	iter.current.Drop()
}

func (iter *forwardIterList[T]) Next() bool {
	iter.current = iter.next

	if iter.next != nil && iter.next != iter.root {
		iter.next = iter.next.next
		return true
	}

	return false
}

func (iter *forwardIterList[T]) Get() T {
	return iter.current.value
}

func (iter *forwardIterList[T]) GetPtr() *T {
	return &iter.current.value
}

type reverseIterList[T any] struct {
	root    *listElement[T]
	prev    *listElement[T]
	current *listElement[T]
}

func (list *List[T]) ReverseIter() IterDrop[T] {
	return &reverseIterList[T]{
		root: &list.root,
		prev: list.root.prev,
	}
}

func (iter *reverseIterList[T]) Drop() {
	iter.current.Drop()
}

func (iter *reverseIterList[T]) Next() bool {
	iter.current = iter.prev

	if iter.prev != nil && iter.current != iter.root {
		iter.prev = iter.prev.prev
		return true
	}

	return false
}

func (iter *reverseIterList[T]) Get() T {
	return iter.current.value
}

func (iter *reverseIterList[T]) GetPtr() *T {
	return &iter.current.value
}
