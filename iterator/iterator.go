package iterator

import (
	"errors"

	"github.com/johnsonlee-debug.com/GenGo/result"
)

type Iterable[T any] interface {
	Iterator() Iterator[T]
	ForEach(func(T))
}

type Iterator[E any] interface {
	HasNext() bool
	Next() result.Result[E]
	Remove() result.Result[result.Unit]
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

func (m *mapIterator[A, B]) Next() result.Result[B] {
	return result.Fmap(m.transform, m.it.Next())
}

func (m *mapIterator[A, B]) Remove() result.Result[result.Unit] {
	return result.Fail[result.Unit](unSupportedOperationException())
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

func (f *filterIterator[A]) Next() result.Result[A] {
	for f.it.HasNext() {
		next := f.it.Next()
		ok := result.Fmap(f.filter, next).ValOrElse(false)
		if ok {
			return next
		}
	}
	return result.Fail[A](noSuchElementException())
}

func (f *filterIterator[A]) Remove() result.Result[result.Unit] {
	return result.Fail[result.Unit](unSupportedOperationException())
}

func Filter[A any](f func(A) bool, it Iterator[A]) Iterator[A] {
	return &filterIterator[A]{it, f}
}
