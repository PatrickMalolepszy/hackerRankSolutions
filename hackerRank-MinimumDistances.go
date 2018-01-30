package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	result := minimumDistances(a)
	fmt.Println(result)
}

func minimumDistances(a []int) int {
	min := math.MaxUint32
	noneFound := true
	for i := 0 ; i < len(a) ; i++ {
		for j := i+1 ; j < len(a) ; j++ {
			if a[i] == a[j] && j - i < min {
				min = j - i
				noneFound = false
			}
		}
	}
	if noneFound {
		return -1
	}
	return min
}