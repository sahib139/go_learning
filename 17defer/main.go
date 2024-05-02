package main

import "fmt"

func main() {
	defer fmt.Println("lets begin learning")
	fmt.Println("defer learning")

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	myFunc()
}

func myFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	nextFunc()
}

func nextFunc() {
	for i := 10; i > 5; i-- {
		defer fmt.Println(i)
	}
}
