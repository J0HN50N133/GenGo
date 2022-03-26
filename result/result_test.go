package result

import (
	"errors"
	"fmt"
	"testing"
)

type MyWriter struct{}

var count = 0

func (m *MyWriter) Write(a string) Result[int] {
	if count == 2 {
		return Fail[int](errors.New("Write too many times"))
	}
	count += 1
	fmt.Println(a)
	return Ok(len(a))
}

func (m *MyWriter) TraditionWrite(a string) (int, error) {
	if count == 2 {
		return -1, errors.New("Write too many times")
	}
	count += 1
	fmt.Println(a)
	return len(a), nil
}

func TraditionTest(t *testing.T) {
	count = 0
	w := MyWriter{}
	wrap := Wrap1(w.TraditionWrite)
	result := wrap("Write 1").
		Then(func() Result[int] { return wrap("Write 2") }).
		Then(func() Result[int] { return wrap("Write 3") }).
		Then(func() Result[int] { return wrap("Write 4") })

	if result.IsFail() {
		fmt.Println(result.ErrorOrNil())
	} else {
		fmt.Println(result.ValOrElse(0o1))
	}
}

func ResultTest(t *testing.T) {
	count = 0
	w := MyWriter{}
	result := w.Write("Write 1").
		Then(func() Result[int] { return w.Write("Write 1") }).
		Then(func() Result[int] { return w.Write("Write 2") }).
		Then(func() Result[int] { return w.Write("Write 3") }).
		Then(func() Result[int] { return w.Write("Write 4") }).
		Then(func() Result[int] { return w.Write("Write 5") })

	if result.IsFail() {
		fmt.Println(result.ErrorOrNil())
		fmt.Println(result.ValOrElse(-1))
	} else {
		fmt.Println(result.ValOrElse(-1))
	}
}

type SomeResource struct {
	DatabaseId int
}

func GetSomeResource(id int) (*SomeResource, error) {
	if id == 1 {
		return &SomeResource{810975}, nil
	} else {
		return nil, errors.New("Resource not found")
	}
}

func GiveMeASafeFunc() {
}

func IfOkTest(t testing.T) {
	safeLogic := func(id int) {
		Wrap1(GetSomeResource)(id).
			IfOk(func(s *SomeResource) { fmt.Println(s.DatabaseId) }).
			IfFail(func(e error) { fmt.Println(e) })
	}
	safeLogic(1)
	safeLogic(2)
}

func FoldTest(t *testing.T) {
	safeLogic := func(id int) Result[*SomeResource] {
		return Wrap1(GetSomeResource)(id).
			Fold(func(s *SomeResource) {},
				func(e error) { fmt.Println(e) })
	}
	safeLogic(1).
		IfOk(func(s *SomeResource) { s.DatabaseId = 9999 }).
		IfOk(func(s *SomeResource) { fmt.Println(s.DatabaseId) })

	safeLogic(2).
		IfOk(func(s *SomeResource) { s.DatabaseId = 2132173 }).
		IfOk(func(s *SomeResource) { fmt.Println(s.DatabaseId) })
}
