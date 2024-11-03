package main

import (
	"fmt"
)

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}

func main() {
	var x = make([]int, 5)
	//printSlice(x)
	var y []int

	if x == nil {
		fmt.Println("slice x is nil")
	} else {
		printSlice(x)
	}
	if y == nil {
		fmt.Println("slice y is nil")
		printSlice(y)
	} else {
		printSlice(y)
	}

}
