package main

import "fmt"

func main() {

	Fruits := []string{"apple", "orange", "banana", "grapes"}

	for i := 0; i < len(Fruits); i++ {
		fmt.Println(Fruits[i])
	}

	for i := range Fruits {
		fmt.Println(Fruits[i])
	}

	for idx, value := range Fruits {
		fmt.Println(idx, value)
	}

}
