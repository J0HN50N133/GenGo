package deque

import (
	"errors"

	"github.com/johnsonlee-debug.com/GenGo/iterator"
	"github.com/johnsonlee-debug.com/GenGo/result"
)

type node[Item any] struct {
	// The value stored with this node
	val *Item
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *node[Item]

	// which deque this node belong to
	deque *Deque[Item]
}

type Deque[Item any] struct {
	iterator.Iterable[Item]
	dummy node[Item]
	len   int
}

type innerIterator[Item any] struct {
	position *node[Item]
}

func noSuchElementException() error {
	return errors.New("No Such Element Exception")
}

func (i *innerIterator[Item]) HasNext() bool {
	return i.position != &i.position.deque.dummy
}

func (i *innerIterator[Item]) Next() result.Result[Item] {
	if !i.HasNext() {
		return result.Fail[Item](noSuchElementException())
	}
	i.position = i.position.next
	return result.Ok(*i.position.prev.val)
}

func (i *innerIterator[Item]) Remove() result.Result[result.Unit] {
	panic("not implemented") // TODO: Implement
}

func (d *Deque[Item]) Iterator() iterator.Iterator[Item] {
	return &innerIterator[Item]{d.dummy.next}
}

func (d *Deque[Item]) ForEach(f func(Item)) {
	it := d.Iterator()
	for it.HasNext() {
		it.Next().IfOk(f)
	}
}

func (d *Deque[Item]) Init() *Deque[Item] {
	d.dummy.next = &d.dummy
	d.dummy.prev = &d.dummy
	d.dummy.deque = d
	d.len = 0
	return d
}

func New[Item any]() *Deque[Item] {
	return new(Deque[Item]).Init()
}

func (d *Deque[Item]) Len() int { return d.len }

func (d *Deque[Item]) IsEmpty() bool { return d.len == 0 }

func (d *Deque[Item]) lazyInit() {
	if d.dummy.next == nil {
		d.Init()
	}
}

func (d *Deque[Item]) insert(e, at *node[Item]) {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.deque = d
	d.len++
}

func (d *Deque[Item]) insertValue(v Item, at *node[Item]) {
	d.insert(&node[Item]{val: &v}, at)
}

func (d *Deque[Item]) PushFront(e Item) {
	d.lazyInit()
	d.insertValue(e, &d.dummy)
}

func (d *Deque[Item]) PushBack(e Item) {
	d.lazyInit()
	d.insertValue(e, d.dummy.prev)
}

func (d *Deque[Item]) remove(e *node[Item]) Item {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = nil
	e.next = nil
	d.len--
	return *e.val
}

func (d *Deque[Item]) PopFront() result.Result[Item] {
	if d.IsEmpty() {
		return result.Fail[Item](errors.New("NoSuchElementException"))
	}
	return result.Ok(d.remove(d.dummy.next))
}

func (d *Deque[Item]) PopBack() result.Result[Item] {
	if d.IsEmpty() {
		return result.Fail[Item](errors.New("NoSuchElementException"))
	}
	return result.Ok(d.remove(d.dummy.prev))
}
