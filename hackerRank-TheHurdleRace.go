package main

import "fmt"

func main() {
	var n, k, h int
	fmt.Scan(&n, &k)
	bevs := 0
	for i := 0 ; i < n ; i++ {
		fmt.Scan(&h)
		for k + bevs < h {
			bevs++
		}
	}
	fmt.Println(bevs)
}