package main

import (
	"fmt"
)

func Addation(num1 int, num2 int) int {
	//完成两个数的相加
	// 大写的函数名称为公共函数，可以被其他的包所调用
	var sum = 0
	sum += num1
	sum += num2
	return sum
	//fmt.Println(sum)
}

func addation(num1 int, num2 int) {
	//完成两个数的相加
	// 小写的函数名为私有函数，只能在本包中被调用
	// 如果没有返回值,则返回值列表可以不用写,但是如果要用return做返回值,则需要编写返回值列表
	var sum = 0
	sum += num1
	sum += num2
	//return sum
	fmt.Println(sum)
}

func fuhe(num1, num2 int) (int, int) {
	// 可以传入两个返回值列表,让函数返回两个返回值,但是在做函数运算的时候,需要用两个变量来接收这个函数的返回值
	// 不能使用一个变量进行接受,但是可以用 _ 来不接受
	var sum1 = 0
	sum1 += num1
	sum1 += num2
	var sum2 = num1 - num2
	return sum1, sum2
}

func main() {
	fmt.Println("---------------start function--------------")
	addation(1, 2)
	sum1 := Addation(1, 3)
	fmt.Println(sum1)
	//Addation(1, 3)
	//addation(1, 2)

	var sum2, sum3 int
	sum2, sum3 = fuhe(1, 2)
	sum4, _ := fuhe(1, 2) //两个返回值只接收其中一个
	fmt.Println(sum2, sum3)
	fmt.Println(sum4)

	fmt.Println("---------------end function--------------")
}
