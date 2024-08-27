package main

import "fmt"

func test(num ...int) {
	// 可以传入多个可变参数,用... 来表示
	// 函数内部处理可变参数的时候,可以讲可变参数当作切片来进行处理
	// 可以讲可变参数进行遍历
	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}

}

func main() {
	test(1, 1, 1, 1, 1, 1, 1, 1)
}
