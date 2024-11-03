package main

import "fmt"

func printslice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}

func main() {
	var num1 []int
	num1 = append(num1, 0)
	num1 = append(num1, 1)
	num1 = append(num1, 2)
	num1 = append(num1, 3)
	fmt.Println(num1)

	// 使用copy的时候需要指定slice的长度和容量，num2的长度和容量需要>= num1 长度容量，否则无法copy
	// copy 方法是建立了slice的副本，两个slice之间没有任何联系
	num2 := make([]int, len(num1), cap(num1))
	copy(num2, num1)
	printslice(num2)

}
