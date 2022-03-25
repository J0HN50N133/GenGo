package deque

import (
	"fmt"
	"testing"

	"github.com/johnsonlee-debug.com/GenGo/iterator"
)

func TestInit(t *testing.T) {
	d := New[int]()
	if d.dummy.next != d.dummy.prev {
		t.Fail()
		t.Errorf("dummy head and prev are not the same")
	}
	if &d.dummy != d.dummy.next {
		t.Fail()
		t.Errorf("dummy next do not point to dummy")
	}
}

func TestIterator(t *testing.T) {
	d := New[int]()
	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)
	d.PushFront(4)
	d.PushFront(5)
	d.PushFront(6)
	d.ForEach(func(i int) {
		t.Log(i)
	})
}

func TestMap(t *testing.T) {
	d := New[int]()
	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)
	d.PushFront(4)
	d.PushFront(5)
	d.PushFront(6)
	it := iterator.Map(func(i int) []int { return []int{i, i + 1, i + 2, i + 3} }, d.Iterator())
	for it.HasNext() {
		it.Next().IfOk(func(i []int) { fmt.Println(i) })
	}
}

func TestFilter(t *testing.T) {
	d := New[int]()
	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)
	d.PushFront(4)
	d.PushFront(5)
	d.PushFront(6)
	it := iterator.Map(func(i int) []int { return []int{i, i + 1, i + 2, i + 3} }, d.Iterator())
	fit := iterator.Filter(func(i []int) bool { return len(i) >= 1 && i[0]%2 == 0 }, it)
	for fit.HasNext() {
		fit.Next().IfOk(func(i []int) { fmt.Println(i) })
	}
}
