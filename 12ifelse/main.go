package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age: ")
	number, _ := reader.ReadString('\n')

	num, err := strconv.Atoi(strings.TrimSpace(number))

	if err != nil {
		fmt.Println(err)
	}
	if num >= 18 {
		fmt.Println("Can Vote")
	} else {
		fmt.Println("Cannot Vote")
	}

	if num := 10; num <= 10 {
		fmt.Println("yes")
	}
}
