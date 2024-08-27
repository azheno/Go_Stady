package main	 

import "fmt" 

func add(num1,num2 int) int { 
	sum := num1 + num2 
	return sum 
}

func main (){
	
	var num1,num2 int  
	fmt.Scan(&num1,&num2) 	
	fmt.Println(add(num1,num2)) 


}
