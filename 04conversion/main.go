package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the rating of our pizza :")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	parseInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The new Rating is :", parseInput+1)
	}
}
