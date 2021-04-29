package main

// soln for exercise 12
import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s [3]string
	s[0] = "1 b 3 1 a a b"
	s[1] = "11 a 1 1 1 1 a a"
	s[2] = "-1 b 1 -4 a 1"

	arguments := os.Args
	thresh, _ := strconv.Atoi(arguments[1])
	counts := make(map[string]int)
	fmt.Println(thresh)
	for i := 0; i < len(s); i++ {
		data := strings.Fields(s[i])
		for _, word := range data {
			_, ok := counts[word]
			if ok {
				counts[word] = counts[word] + 1
			} else {
				counts[word] = 1
			}
		}
	}
	for key, _ := range counts {
		if counts[key] > thresh {
			fmt.Printf("%s -> %d \n", key, counts[key])
		}
	}
}
