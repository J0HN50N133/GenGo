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
		return Result[int]{-1, errors.New("Write too many times")}
	}
	count += 1
	fmt.Println(a)
	return Result[int]{count, nil}
}

func ResultTest() {
	w := MyWriter{}
	result := w.Write("Write 1").
		Then(func() Result[int] { return w.Write("Write 1") }).
		Then(func() Result[int] { return w.Write("Write 2") }).
		Then(func() Result[int] { return w.Write("Write 3") }).
		Then(func() Result[int] { return w.Write("Write 4") }).
		Then(func() Result[int] { return w.Write("Write 5") })
	if result.Error() != nil {
		fmt.Println(result.Error())
	} else {
		fmt.Println(result.Value())
	}
}

func main() {
	ResultTest()
}
