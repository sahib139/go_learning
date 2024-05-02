package main

import (
	"fmt"
	"sort"
)

func main() {
	var dynamicArr []int
	fmt.Println(dynamicArr, len(dynamicArr))
	for i := 0; i < 5; i++ {
		dynamicArr = append(dynamicArr, i)
	}
	dynamicArr = append(dynamicArr, 10)
	fmt.Println(dynamicArr, len(dynamicArr), dynamicArr[2:3])

	score := make([]int, 3)
	score[0] = 50
	score[1] = 20
	score[2] = 100
	score = append(score, 0)
	fmt.Println(score)
	sort.Ints(score)
	fmt.Println(score, sort.IntsAreSorted(score))

	score = append(score[:2], score[(2+1):]...)
	fmt.Println(score)
}
