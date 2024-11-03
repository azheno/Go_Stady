package main

import (
	"fmt"
)

func main() {
	// 创建map
	map1 := make(map[string]int)
	map1["A"] = 1
	map1["B"] = 2
	map1["C"] = 3
	map1["D"] = 4
	map1["E"] = 5
	map1["F"] = 6
	map1["G"] = 7
	map1["H"] = 8
	map1["I"] = 9

	// 使用for 循环和range遍历map，map存储值是无序的
	for x := range map1 {
		fmt.Println(x, ":", map1[x])
	}
	fmt.Println("map1 长度为： ", len(map1))
	// delete() 删除map key value
	delete(map1, "A")
	println("------------------------")
	for y := range map1 {
		fmt.Println(y, ":", map1[y])
	}

	// len 确定map长度
	fmt.Println("map1 长度为： ", len(map1))
}
