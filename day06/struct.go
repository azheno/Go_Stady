package main

import "fmt"

func main() {
	type message struct {
		name string
		age  int
		sex  string
	}

	var user1 message
	user1.name = "wangzhendong"
	user1.age = 14
	user1.sex = "man"

	fmt.Print(user1)

}
