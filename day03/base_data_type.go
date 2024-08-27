package main	 

import (
	"fmt" 
	"unsafe" 
)


func main(){

	// int 类型
	var num1 int = -1212414141421241
	fmt.Println("num1 = ",num1) 
	fmt.Println("----------------------------------------")

	// uint 类型 
	var num2 uint = 1123123123 
	//var num2 uint = -12313123 // uint不能使用负数 
	fmt.Println("num2 = ",num2) 
	fmt.Println("----------------------------------------")
	
	// 计算变量占用字节大小 unsafe
	fmt.Println("num2变量占用字节大小 = ",unsafe.Sizeof(num2)) 
	fmt.Println("----------------------------------------")

	// float 类型 (可以接收正负以及科学计数法) 
	var num3 float32 = 12.12 
	fmt.Println("num3 = ",num3) 
	fmt.Println("----------------------------------------")

	// 

}

