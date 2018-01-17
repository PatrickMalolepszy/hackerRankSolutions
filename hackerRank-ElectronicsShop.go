package main

import "fmt"

func main() {
	var s, n, m int
	fmt.Scan(&s, &n, &m)
	keyboards := make([]int, n)
	for i := 0 ; i < n ; i++ {
		fmt.Scan(&keyboards[i])
	}
	mice := make([]int, m)
	for i := 0 ; i < m ; i++ {
		fmt.Scan(&mice[i])
	}
	result := electronicsShop(s, keyboards, mice)
	fmt.Println(result)
}

func electronicsShop(s int, keyboards, mice []int) int {
	max := -1
	for i := range keyboards {
		for j := range mice {
			price := keyboards[i] + mice[j]
			if price > max && price <= s {
				max = price
			}
		}
	}
	return max
}