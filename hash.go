package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node

	Size int
}

func hashFunction(i, size int) int {
	return (i % size)
}

func hashIsIn(hash *HashTable, value int) bool {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				if t.Value == value {
					return true
				}
				t = t.Next
			}
		}
	}
	return false
}

func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	if hashIsIn(hash, value) {
		fmt.Println("Value already in table")
		return 0
	}
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverse(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func main() {
	table := make(map[int]*Node, 10)
	hash := &HashTable{Table: table, Size: 10}
	fmt.Println("Number of space:", hash.Size)
	for i := 0; i < 95; i++ {
		insert(hash, i)
	}
	traverse(hash)
}
