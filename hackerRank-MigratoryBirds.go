package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n)
	birds := make([]int, 5)
	for i := 0 ; i < n ; i++ {
		fmt.Scan(&t)
		birds[t-1]++
	}
	max := 0
	birdType := -1
	for i := 4 ; i >= 0 ; i-- {
		if birds[i] >= max {
			max = birds[i]
			birdType = i+1
		}
	}
	fmt.Println(birdType)
}