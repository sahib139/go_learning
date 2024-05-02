package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time")

	presentTime := time.Now()
	currentTime := presentTime.Format("02-01-2006 Monday 15:04:05")
	fmt.Println(presentTime, currentTime)
	createDate := time.Date(2025, time.February, 23, 1, 55, 0, 0, time.UTC)
	fmt.Println(createDate.Format("02-01-2006 Monday"))
}
