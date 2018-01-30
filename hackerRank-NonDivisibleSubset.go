package main

import (
	"fmt"
)

func main() {
	var n, k, t int
	fmt.Scan(&n, &k)
	divisibleCount := make(map[int]int)
	set := make(map[int]bool)
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		if !set[t] {
			divisibleCount[t%k]++ // need to keep track of the number of unique numbers which are divisible by k
			set[t] = true
		}
	}
	result := nonDivisibleSubset(k, divisibleCount)
	fmt.Println(result)
}

func nonDivisibleSubset(k int, set map[int]int) int {
	count := 0
	if set[0] > 0 {
		count++
	}
	for i, j := 1, k-1; i < j; i, j = i+1, j-1 {
		if set[i] > set[j] {
			count += set[i]
		} else {
			count += set[j]
		}
	}
	if k%2 == 0 {
		count ++
	}
	return count
}
