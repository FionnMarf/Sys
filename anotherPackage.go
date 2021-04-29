package anotherPackage

import "fmt"

const Pi = "3.14159"  // Upper case therefore a public variable
const version = "1.1" // lower case therefore a private variable

// Runs if package is used, might be needed to open db or net connects
func init() {
	fmt.Println("The init function of anotherPackage")
}

func Add(x, y int) int {
	return x + y
}

func Println(x int) {
	fmt.Println(x)
}

func Version() {
	fmt.Println("The version of the package is", version)
}
