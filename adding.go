package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Printfln("1 + 2 is %d", add(1, 2))
}

func add(a, b int) int {
	return a + b
}
