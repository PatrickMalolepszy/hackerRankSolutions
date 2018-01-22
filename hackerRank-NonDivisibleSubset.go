package main

import (
	"fmt"
)

func main() {
	var n, k, t int
	fmt.Scan(&n, &k)
	set := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		set[t%k]++
	}
	result := nonDivisibleSubset(n, k, set)
	fmt.Println(result)
}

func nonDivisibleSubset(n, k int, set map[int]int) int {
	count := 0
	//fmt.Println(set)
	count += set[0]
	for i, j := 1, k-1; i < j; i, j = i+1, j-1 {
		if set[i] > set[j] {
			count += set[i]
		} else {
			count += set[j]
		}
	}
	if k%2 == 0 && k > 2 {
		count += set[k/2]
	}
	return count
}
