package main

import "fmt"

func main() {
	var d, n, m int
	fmt.Scan(&n)
	bar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&bar[i])
	}
	fmt.Scan(&d, &m)
	result := birthdayChocolate(d, m , bar)
	fmt.Println(result)
}

func birthdayChocolate(d int, m int, bar []int) int {
	sum := 0
	count := 0
	for i := 0 ; i < m ; i++ {
		sum += bar[i]
	}
	if sum == d {
		count++
	}
	for i := m ; i < len(bar) ; i++ {
		sum -= bar[i-m]
		sum += bar[i]
		if sum == d {
			count++
		}
	}
	return count
}
