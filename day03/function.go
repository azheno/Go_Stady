package main

import (
	"fmt"
)

func addation(num1 int, num2 int) int {
	//完成两个数的相加
	var sum = 0
	sum += num1
	sum += num2
	return sum
}

func main() {
	fmt.Println("---------------start function--------------")
	sum := addation(1, 2)
	fmt.Println(sum)
	fmt.Println("---------------end function--------------")
}
