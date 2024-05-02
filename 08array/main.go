package main

import "fmt"

func main() {
	var arr [3]string
	arr[0] = "1"
	arr[2] = "2"
	fmt.Println(arr)

	var numArr = [2]string{"sahib", "singh"}
	fmt.Println(numArr, len(numArr))
}
