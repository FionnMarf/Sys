package main

import "fmt"

// arrays are old school and slow af, static and need to be copied to pass them around

func main() {
	myArray := [4]int{1, 2, 4, -4}

	twoD := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	threeD := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	fmt.Println(myArray[0])
	twoD[1][2] = 15
	threeD[0][1][1] = -1
}
