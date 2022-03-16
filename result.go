package main

type Result[A any] struct {
	result A
	err    error
}

func (r Result[A]) Value() A {
	return r.result
}

func (r Result[A]) Then(f func() Result[A]) Result[A] {
	if r.Error() == nil {
		return f()
	} else {
		return r
	}
}

func (r Result[A]) Error() error {
	return r.err
}
