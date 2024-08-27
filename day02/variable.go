package main	 

import "fmt" 

func main(){

	// 第一种声明变量的方法 (声明后赋值)
	fmt.Println("----------------------------------------")
	var type1 int 
	fmt.Println("不赋值时type1默认值 = ",type1)
	type1 = 18 
	fmt.Println("type1 = ",type1) 
	fmt.Println("----------------------------------------")

	// 第二种声明变量的方法 (声明时赋值)
	var type2 int = 18
	fmt.Println("type2 = ",type2) 
	fmt.Println("----------------------------------------") 

	// 第三种声明变量的方法 (不声明类型，自动类型推断)
	var type3 = 18 
	fmt.Println("type3 = ",type3)
	fmt.Println("----------------------------------------")

	// 第四种声明变量的方法 (省略关键字声明，需要使用:= )
	type4 := 18 
	fmt.Println("type4 = ",type4)
	fmt.Println("----------------------------------------")

	// 第五种变量声明的方法 (多变量声明,声明后赋值) 
	var type5,type6 int 
	type5 = 18 
	type6 = 18 
	fmt.Println("type5 = ",type5,"\ntype6 = ",type6)  
	fmt.Println("----------------------------------------")
	
	// 第六种变量声明方法 (多变量声明，声明时赋值) 
	var type7,type8 = 18,18
	fmt.Println("type7 = ",type7,"\ntype8 = ",type8)
	fmt.Println("----------------------------------------")

	// 第七种变量声明方法 (多变量声明，省略关键字声明，需要使用:=) 
	type9,type10 := 18,18 
	fmt.Println("type9 = ",type9,"\ntype10 = ",type10)
	fmt.Println("----------------------------------------")
	
	// 第八种变量声明方法 (一次性声明) 
	var (
		type11 = 18  
		type12 = 18 
	)
	fmt.Println("type11 = ",type11,"\ntype12 = ",type12)  
	fmt.Println("----------------------------------------")
}
