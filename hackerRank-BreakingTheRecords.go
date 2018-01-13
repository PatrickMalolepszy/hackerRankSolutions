package main

import "fmt"

func main() {
	var n, k uint32
	fmt.Scan(&n)
	fmt.Scan(&k)
	min, max := k, k
	minCount, maxCount := 0, 0
	for i := uint32(0); i < n-1; i++ {
		fmt.Scan(&k)
		if k < min {
			min = k
			minCount++
		}
		if k > max {
			max = k
			maxCount++
		}
	}
	fmt.Println(maxCount, minCount)
}
