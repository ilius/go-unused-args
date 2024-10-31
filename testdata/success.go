package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("Hello world!")
}

func TestFunc(t *testing.T) {}

func FuncVars(x int) {
	do := func() {
		fmt.Println(x)
	}
	do()
}
