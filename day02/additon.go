package main	 

import "fmt" 

func add(num1,num2 int) int {
	var sum int 
	sum = num1 + num2 
	return sum 
}

func main (){
	
	var num1,num2,sum int  

	fmt.Scan(&num1,&num2) 

	sum = add(num1,num2)  
	
	fmt.Printf("%d和%d的和为%d\n",num1,num2,sum) 

}
