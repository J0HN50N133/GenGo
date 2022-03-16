package main

import (
	"errors"
	"fmt"
)

type MyWriter struct {
}

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
func TraditionTest() {
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
		fmt.Println(result.ValOrElse(01))
	}
}
func ResultTest() {
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
	} else {
		fmt.Println(result.ValOrElse(01))
	}
}

func main() {
	TraditionTest()
	ResultTest()
}
