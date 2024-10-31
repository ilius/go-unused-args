package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("Hello world!")
}

func FuncVars(x int) {
	do := func() {
		fmt.Println(x)
	}
	do()
}

func newTypePair[K any, V any]() {
}

func testIndexListExpr() {
	newTypePair[int, int]
}

func TestFunc(t *testing.T) {}
