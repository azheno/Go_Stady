package main

import (
	"fmt"
)

func main() {

	// 第一种声明变量的方法 (声明后赋值)
	fmt.Println("----------------------------------------")
	var type1 int
	fmt.Println("不赋值时type1默认值 = ", type1)
	type1 = 18
	fmt.Println("type1 = ", type1)
	fmt.Println("----------------------------------------")

	// 第二种声明变量的方法 (声明时赋值)
	var type2 int = 18
	fmt.Println("type2 = ", type2)
	fmt.Println("----------------------------------------")

	// 第三种声明变量的方法 (不声明类型，自动类型推断)
	var type3 = 18
	fmt.Println("type3 = ", type3)
	fmt.Println("----------------------------------------")

	// 第四种声明变量的方法 (省略关键字声明，需要使用:= )
	type4 := 18
	fmt.Println("type4 = ", type4)
	fmt.Println("----------------------------------------")

	// 第五种变量声明的方法 (多变量声明,声明后赋值)
	var type5, type6 int
	type5 = 18
	type6 = 18
	fmt.Println("type5 = ", type5, "\ntype6 = ", type6)
	fmt.Println("----------------------------------------")

	// 第六种变量声明方法 (多变量声明，声明时赋值)
	var type7, type8 = 18, 18
	fmt.Println("type7 = ", type7, "\ntype8 = ", type8)
	fmt.Println("----------------------------------------")

	// 第七种变量声明方法 (多变量声明，省略关键字声明，需要使用:=)
	type9, type10 := 18, 18
	fmt.Println("type9 = ", type9, "\ntype10 = ", type10)
	fmt.Println("----------------------------------------")

	// 第八种变量声明方法 (一次性声明)
	var (
		type11 = 18
		type12 = 18
	)
	fmt.Println("type11 = ", type11, "\ntype12 = ", type12)
	fmt.Println("----------------------------------------")

	// 常量声明多重赋值
	const a, b, c = 1, false, "abc"
	fmt.Println(a, b, c)

	// 枚举常量
	// 常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
	const (
		x1 uint16 = 16
		y
		s = "abc"
		z
	)
	fmt.Printf("%T,%v\n", y, y)
	fmt.Printf("%T,%v\n", z, z)

	// iota，特殊常量，可以认为是一个可以被编译器修改的常量
	// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
	// 自增默认是int类型，可以自行进行显示指定类型
	const (
		aa = iota //0
		bb        //1
		cc        //2
		dd = "ha" //独立值，iota += 1
		ee        //"ha"   iota += 1
		ff = 100  //iota +=1
		gg        //100  iota +=1
		hh = iota //7,恢复计数
		ii        //8
	)
	fmt.Println(aa, bb, cc, dd, ee, ff, gg, hh, ii)

	//fmt.Println("请输入一个字符串：")
	//reader := bufio.NewReader(os.Stdin) // 创建读取对象
	//s1, _ := reader.ReadString('\n')    // 读取存入的字符串
	//fmt.Println("读到的数据：", s1)

	var x interface{} // 指定一个数据类型为空接口 interface{} 用于接收任何类型的值
	x = 123

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}

}
