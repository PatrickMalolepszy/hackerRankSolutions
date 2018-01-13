package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
	}
	result := divisibleSumPairs(k, c)
	fmt.Println(result)
}

func divisibleSumPairs(k int, c []int) int {
	count := 0
	for i := 0 ; i < len(c) - 1 ; i++ {
		for j := i+1 ; j < len(c) && i < j ; j++ {
			if (c[i] + c[j]) % k == 0 {
				count++
			}
		}
	}
	return count
}
