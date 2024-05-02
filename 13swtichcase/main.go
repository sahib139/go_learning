package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("piece opened")
	case 2:
		fmt.Println("piece not opened")
	case 3:
		fmt.Println("piece not opened")
		fallthrough
	case 4:
		fmt.Println("piece not opened")
	case 5:
		fmt.Println("piece not opened")
	case 6:
		fmt.Println("piece opened and player can play again!")
	}
}
