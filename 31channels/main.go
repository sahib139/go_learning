package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Learning about channels")

	mych := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(mych <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		val, isChannelOpen := <-mych
		if isChannelOpen {
			fmt.Println("value from other go routine is ", val)
		} else {
			fmt.Println("Channel is close")
		}

	}(mych, wg)

	go func(mych chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		mych <- 5
		close(mych)
	}(mych, wg)

	wg.Wait()
}
