package iterator

import (
	"errors"

	. "github.com/johnsonlee-debug.com/GenGo/result"
)

type Iterable[T any] interface {
	Iterator() Iterator[T]
	ForEach(func(T))
}

type Iterator[E any] interface {
	HasNext() bool
	Next() Result[E]
	Remove() Result[Unit]
}

func noSuchElementException() error {
	return errors.New("No Such Element Exception")
}

func unSupportedOperationException() error {
	return errors.New("Unsupported Operation Exception")
}

type mapIterator[A, B any] struct {
	it        Iterator[A]
	transform func(A) B
}

func (m *mapIterator[A, B]) HasNext() bool {
	return m.it.HasNext()
}

func (m *mapIterator[A, B]) Next() Result[B] {
	return Fmap(m.transform, m.it.Next())
}

func (m *mapIterator[A, B]) Remove() Result[Unit] {
	return Fail[Unit](unSupportedOperationException())
}

func ForEach[A any](f func(A), it Iterator[A]) {
	for it.HasNext() {
		it.Next().IfOk(f)
	}
}

func Map[A, B any](f func(A) B, it Iterator[A]) Iterator[B] {
	return &mapIterator[A, B]{it, f}
}

type filterIterator[A any] struct {
	it     Iterator[A]
	filter func(A) bool
}

func (f *filterIterator[A]) HasNext() bool {
	return f.it.HasNext()
}

func (f *filterIterator[A]) Next() Result[A] {
	for f.it.HasNext() {
		next := f.it.Next()
		ok := Fmap(f.filter, next).ValOrElse(false)
		if ok {
			return next
		}
	}
	return Fail[A](noSuchElementException())
}

func (f *filterIterator[A]) Remove() Result[Unit] {
	return Fail[Unit](unSupportedOperationException())
}

func Filter[A any](f func(A) bool, it Iterator[A]) Iterator[A] {
	return &filterIterator[A]{it, f}
}
