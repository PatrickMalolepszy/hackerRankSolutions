package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	canSort, swap, i, j := almostSorted(n, a)
	if canSort {
		fmt.Println("yes")
		if swap {
			fmt.Println("swap", i, j)
		} else {
			fmt.Println("reverse", i, j)
		}
	} else {
		fmt.Println("no")
	}
}

func almostSorted(n int, a []int) (bool, bool, int, int) {
	b := make([]int, n)
	copy(b, a)
	sort.Ints(b)
	swapCount := 0
	startR := -1
	endR := -1
	//fmt.Println(a)
	//fmt.Println(b)
	for i := range a {
		if a[i] != b[i] {
			swapCount++
			if startR == -1 {
				startR = i
			} else {
				endR = i
			}
		}
	}
	if swapCount == 2 {
		return true, true, startR + 1, endR + 1
	}
	canReverse := true
	for i := endR; i > startR; i-- {
		if a[i] > a[i-1] {
			canReverse = false
			break
		}
	}
	if canReverse {
		return true, false, startR + 1, endR + 1
	}
	return false, false, -1, -1
}
