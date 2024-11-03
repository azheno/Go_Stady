package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// int 类型
	var num1 int = -1212414141421241
	fmt.Println("num1 = ", num1)
	fmt.Println("----------------------------------------")

	// uint 类型
	var num2 uint = 1123123123
	//var num2 uint = -12313123 // uint不能使用负数
	fmt.Println("num2 = ", num2)
	fmt.Println("----------------------------------------")

	// 计算变量占用字节大小 unsafe
	fmt.Println("num2变量占用字节大小 = ", unsafe.Sizeof(num2))
	fmt.Println("----------------------------------------")

	// float 类型 (可以接收正负以及科学计数法)
	var num3 float32 = 12.12
	fmt.Println("num3 = ", num3)
	fmt.Println("----------------------------------------")

	// 数组

	// 数组赋值与遍历
	var n [2]int
	for i := 0; i < len(n); i++ {
		n[i] = i
	}
	for j := 0; j < len(n); j++ {
		fmt.Printf("n[%d] = %d \n", j, n[j])
	}

	// 数组定义
	var n1 = []int{1, 2, 3, 4, 5}
	fmt.Println(n1)

	// 其他定义方式
	n2 := []int{1, 2, 3}
	for i := 0; i < len(n2); i++ {
		fmt.Println(n2[i])
	}

	a := []int{1, 2, 3, 4, 5, 6, 32, 14, 1324, 12, 431, 243, 23, 41, 431, 2412}
	fmt.Println("a数组中有：", len(a), "个元素")
	// [] 和 [...] 是等价关系
	a1 := []int{12, 3, 12, 3, 123, 1, 23, 123, 13}
	fmt.Println(len(a1))

	// 使用rang遍历数组

	a2 := []int{1, 2, 13, 1, 23, 123, 21, 31, 23, 12}
	a3 := int(0)
	for i, v := range a2 {
		fmt.Println(i, v)
		a3 += v
	}
	fmt.Println(a3)

	// slice
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("slice1: %d\n", slice1)
	var slice2 []int = slice1[1:4]
	fmt.Println("slice2:", slice2)
	// len() 函数用来测量切片长度
	long1 := len(slice1)
	fmt.Println(long1)

	// 通过make创建切片
	slice3 := make([]int, 4)
	fmt.Printf("slice3: %d\n", slice3)

	// cap() 函数用来测量切片的容量
	// 切片的容量是从创建切片的索引开始的底层数组中元素的数量
	fmt.Printf("slice3 cap: %d\n", cap(slice3))
	fmt.Printf("slice3 len: %d\n", len(slice3))

}
