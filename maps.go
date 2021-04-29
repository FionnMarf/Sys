package main

import "fmt"

func main() {
	aMap := make(map[string]int)
	aMap["Mon"] = 0
	aMap["Tue"] = 1
	aMap["Wed"] = 2
	aMap["Thu"] = 3
	aMap["Fri"] = 4
	aMap["Sat"] = 5
	aMap["Sun"] = 6

	fmt.Printf("Sunday is the %dth day of the week.\n", aMap["Sun"])

	// key Map feature is for checking things exist
	_, ok := aMap["Tuesday"]
	if ok {
		fmt.Printf("The Tuesday Key Exists!\n")
	} else {
		fmt.Printf("The Tuesday Key Does Not Exist!\n")
	}

	count := 0
	for key, _ := range aMap {
		count++
		fmt.Printf("%s", key)
	}
	fmt.Printf("\n")
	fmt.Printf("The aMap has %d elements\n", count)

	anotherMap := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
	}
	anotherMap["Five"] = 5
	count = 0
	for _, _ = range anotherMap {
		count++
	}
	fmt.Printf("anotherMap has %d elements\n", count)

}
