package main

import "fmt"

// 本案例可以作为内存分析, exchangeNum函数中的变量不能跟main函数中的变量交互

func exchangeNum(num1, num2 int) {
	//进行两个数字交换
	var t int
	t = num1
	num1 = num2
	num2 = t
}

func main() {
	var num1, num2 int
	num1 = 10
	num2 = 20
	fmt.Println("转换前的数为:", num1, num2)
	exchangeNum(num1, num2)
	fmt.Println("转换后的数字为:", num1, num2)

}
