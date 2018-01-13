package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&k)
		if k < 38 {
			fmt.Println(k)
			continue
		}
		nextMultiple := 5 * (k/5 + 1)
		difference := nextMultiple - k
		if difference < 3 {
			fmt.Println(nextMultiple)
		} else {
			fmt.Println(k)
		}
	}
}
