package main

import "fmt"

func main() {
	var s, t, a, b, m, n, d int
	fmt.Scan(&s, &t, &a, &b, &m, &n)
	apples, oranges := 0, 0
	for i := 0; i < m; i++ {
		fmt.Scan(&d)
		result := a + d
		if s <= result && result <= t {
			apples++
		}
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&d)
		result := b + d
		if s <= result && result <= t {
			oranges++
		}
	}
	fmt.Println(apples)
	fmt.Println(oranges)
}
