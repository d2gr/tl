package tl

import (
	"testing"

	"github.com/d2gr/tl"
	"github.com/d2gr/tl/iter"
)

func TestList(t *testing.T) {
	var list tl.List[int]

	type expectedValue struct {
		fn       func(int)
		iter     func() tl.IterDrop[int]
		expected []int
	}

	values := iter.ToSlice(iter.Range(0, 10, 1))

	es := []expectedValue{
		{
			fn:       list.PushBack,
			iter:     list.ForwardIter,
			expected: iter.ToSlice(iter.Range(0, 10, 1)),
		}, {
			fn:       list.PushFront,
			iter:     list.ForwardIter,
			expected: iter.ToSlice(iter.Range(9, -1, -1)),
		}, {
			fn:       list.PushBack,
			iter:     list.ReverseIter,
			expected: iter.ToSlice(iter.Range(9, -1, -1)),
		}, {
			fn:       list.PushFront,
			iter:     list.ReverseIter,
			expected: iter.ToSlice(iter.Range(0, 10, 1)),
		},
	}

	for idx, e := range es {
		for _, v := range values {
			e.fn(v)
		}

		iter := e.iter()
		for i := 0; i < len(e.expected); i++ {
			iter.Next()

			if e.expected[i] != iter.Get() {
				t.Fatalf("unexpected value on %d,%d: %d <> %d", idx, i, e.expected[i], iter.Get())
			}
		}

		if iter.Next() {
			t.Fatal("unexpected")
		}

		list.Reset()
	}
}

func TestListPop(t *testing.T) {
	var list tl.List[int]

	type expectedValue struct {
		fn       func(int)
		iter     func() tl.Iter[int]
		expected []int
	}

	iter := list.ForwardIter()
	if iter.Next() {
		t.Fatal("unexpected")
	}

	iter = list.ReverseIter()
	if iter.Next() {
		t.Fatal("unexpected")
	}

	{
		list.PushBack(2)
		list.PushFront(4)
		list.PopFront()
		list.PopBack()
		list.PushBack(3)
		list.PopFront()
		list.PushFront(2)
		list.PushFront(4)
		list.PopBack()
		list.PopFront()

		iter = list.ForwardIter()
		if iter.Next() {
			t.Fatal("unexpected")
		}

		iter = list.ReverseIter()
		if iter.Next() {
			t.Fatal("unexpected")
		}
	}
	{
		list.PushFront(2)
		list.PopBack()

		iter := list.ForwardIter()
		if iter.Next() {
			t.Fatal("unexpected")
		}

		iter = list.ReverseIter()
		if iter.Next() {
			t.Fatal("unexpected")
		}
	}
}

func TestIterDrop(t *testing.T) {
	var list tl.List[int]

	values := iter.ToSlice(iter.Range(1, 11, 1))

	for _, v := range values {
		list.PushBack(v)
	}

	it := list.ForwardIter()

	it.Next()
	values = values[1:]

	for it.Next() {
		it.Drop()

		if it.Get() != values[0] {
			t.Fatalf("got %d, expected %d", it.Get(), values[0])
		}

		values = values[1:]
	}

	if len(values) != 0 {
		t.Fatalf("unexpected: %d", len(values))
	}

	if it.Next() {
		t.Fatalf("unexpected: %d", it.Get())
	}
}

func TestIterDropReverse(t *testing.T) {
	var list tl.List[int]

	values := iter.ToSlice(iter.Range(1, 11, 1))

	for _, v := range values {
		list.PushBack(v)
	}

	values = iter.ToSlice(iter.Range(10, 0, -1))

	it := list.ReverseIter()

	it.Next()
	values = values[1:]

	for it.Next() {
		it.Drop()

		if it.Get() != values[0] {
			t.Fatalf("got %d, expected %d", it.Get(), values[0])
		}

		values = values[1:]
	}

	if len(values) != 0 {
		t.Fatalf("unexpected: %d", len(values))
	}

	if it.Next() {
		t.Fatalf("unexpected: %d", it.Get())
	}
}
