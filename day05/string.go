package main

import "fmt"

func main() {
	name := "wangzhendong"
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c ", name[i])
	}
	for i := 0; i < len(name); i++ {
		fmt.Printf("%d ", name[i])
	}
}
