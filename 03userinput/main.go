package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Rate me out of 10")

	reader := bufio.NewReader(os.Stdin)

	//comma ok || err err
	input, _ := reader.ReadString('\n') // similar to try catch
	fmt.Println("The rating you given me is :", input)
	fmt.Printf("The type of input variable is %T", input)
}
