package result

import (
	"fmt"
	"testing"
)

func doSomething() Unit {
	fmt.Println("doSomething")
	return Unit{}
}

func TestUnit(t *testing.T) {
	result := doSomething()
	fmt.Println(result)
}
