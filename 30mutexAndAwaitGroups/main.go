package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = &sync.WaitGroup{}
	var mut = &sync.RWMutex{}

	var score = []int{0}

	wg.Add(3)
	go func(wt *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("func one")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
	}(wg, mut)

	go func(wt *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("func Two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
	}(wg, mut)

	go func(wt *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("func Three")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
