package main

import (
	"fmt"
)

// NOTES
// Deferred Fn's are resolved in a
// Last In First Out manner

// Will output 2 1 0
func a1() {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")
	}
}

// Will output 3 3 3
func a2() {
	for i := 0; i < 3; i++ {
		defer func() { fmt.Print(i, " ") }()
	}
}

// Will output 2 1 0
func a3() {
	for i := 0; i < 3; i++ {
		defer func(n int) { fmt.Print(n, " ") }(i)
	}
}

func main() {
	a1()
	fmt.Println()
	a2()
	fmt.Println()
	a3()
	fmt.Println()
}
