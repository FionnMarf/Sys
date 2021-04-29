package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStructure, 0)
	a := aStructure{"Fionn", 170, 70}

	mySlice = append(mySlice, a)
	a = aStructure{"Alec", 200, 90}
	mySlice = append(mySlice, a)
	a = aStructure{"Rawat", 160, 55}
	mySlice = append(mySlice, a)
	a = aStructure{"Rory", 185, 80}
	mySlice = append(mySlice, a)

	fmt.Println("0:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight < mySlice[j].weight
	})

	fmt.Println("<:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight > mySlice[j].weight
	})
	fmt.Println(">:", mySlice)
}
