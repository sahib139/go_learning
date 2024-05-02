package main

import "fmt"

func main() {
	fmt.Println("Welcome to go pointer")

	var myPointer *int
	fmt.Println(myPointer)

	num := 10
	var ptr = &num
	fmt.Println(ptr, num)
	*ptr = *ptr * 2
	fmt.Println(*ptr, num)
}
