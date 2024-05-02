package main

import "fmt"

func main() {
	fmt.Println("This module is to learn function!")

	result1 := adder(2, 3)
	fmt.Println(result1)

	result2 := advanceAdder(1, 2, 3, 4, 5)
	fmt.Println(result2)
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func advanceAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
