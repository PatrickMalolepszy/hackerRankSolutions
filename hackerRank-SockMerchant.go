package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n)
	socks := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		socks[t]++
	}
	result := sockMerchant(socks)
	fmt.Println(result)
}

func sockMerchant(socks map[int]int) int {
	count := 0
	for k := range socks {
		count += socks[k] / 2
	}
	return count
}
