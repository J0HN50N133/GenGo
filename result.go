package main

type Result[T any] interface {
	IsOk() bool
	IsFail() bool

	ValOrElse(defaultVal T) T
	PtrOrNil() *T

	Then(func() Result[T]) Result[T]

	ErrorOrNil() error

	IfOk(do func(T)) Result[T]
	OnErr(do func(error)) Result[T]

	Fold(ifok func(T), onerr func(error)) Result[T]
}

type ok[T any] struct {
	val T
	Result[T]
}

func Ok[T any](val T) Result[T] {
	result := new(ok[T])
	result.val = val
	return *result
}
func (o ok[T]) IsOk() bool {
	return true
}

func (o ok[T]) IsFail() bool {
	return false
}

func (o ok[T]) ValOrElse(defaultVal T) T {
	return o.val
}

func (o ok[T]) PtrOrNil() *T {
	return &o.val
}

func (o ok[T]) Then(f func() Result[T]) Result[T] {
	return f()
}

func (o ok[T]) ErrorOrNil() error {
	return nil
}

func (o ok[T]) IfOk(do func(T)) Result[T] {
	do(o.val)
	return o
}

func (o ok[T]) OnErr(_ func(error)) Result[T] {
	return o
}

func (o ok[T]) Fold(do func(T), _ func(error)) Result[T] {
	do(o.val)
	return o
}

type fail[T any] struct {
	err error
	Result[T]
}

func (f fail[T]) IsOk() bool {
	return false
}
func Fail[T any](err error) Result[T] {
	fail := new(fail[T])
	fail.err = err
	return *fail
}

func (f fail[T]) IsFail() bool {
	return true
}

func (f fail[T]) ValOrElse(defaultVal T) T {
	return defaultVal
}

func (f fail[T]) PtrOrNil() *T {
	return nil
}

func (f fail[T]) Then(do func() Result[T]) Result[T] {
	return do()
}

func (f fail[T]) ErrorOrNil() error {
	return f.err
}

func (f fail[T]) IfOk(_ func(T)) Result[T] {
	return f
}

func (f fail[T]) OnErr(do func(error)) Result[T] {
	do(f.err)
	return f
}

func (f fail[T]) Fold(_ func(T), do func(error)) Result[T] {
	do(f.err)
	return f
}

func Fmap[A, B any](f func(a A) B, r Result[A]) Result[B] {
	if r.IsOk() {
		// not elegant
		return Ok[B](f(r.(ok[A]).val))
	} else {
		return Fail[B](r.ErrorOrNil())
	}
}

// Since method could not have type parameter, I have to write ThenDo here
// >>=
func ThenDo[A, B any](r Result[A], do func(A) Result[B]) Result[B] {
	if r.IsOk() {
		return do(r.(ok[A]).val)
	} else {
		return Fail[B](r.ErrorOrNil())
	}
}

func Wrap[T any](f func() (T, error)) func() Result[T] {
	return func() Result[T] {
		val, err := f()
		if err != nil {
			return Fail[T](err)
		} else {
			return Ok[T](val)
		}
	}
}

func Wrap1[A, T any](f func(A) (T, error)) func(A) Result[T] {
	return func(arg A) Result[T] {
		val, err := f(arg)
		if err != nil {
			return Fail[T](err)
		} else {
			return Ok[T](val)
		}
	}
}

func Wrap2[A, B, T any](f func(A, B) (T, error)) func(A, B) Result[T] {
	return func(arg1 A, arg2 B) Result[T] {
		val, err := f(arg1, arg2)
		if err != nil {
			return Fail[T](err)
		} else {
			return Ok[T](val)
		}
	}
}

func Wrap3[A, B, C, T any](f func(A, B, C) (T, error)) func(A, B, C) Result[T] {
	return func(arg1 A, arg2 B, arg3 C) Result[T] {
		val, err := f(arg1, arg2, arg3)
		if err != nil {
			return Fail[T](err)
		} else {
			return Ok[T](val)
		}
	}
}
