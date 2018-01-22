package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	result := viralAdverts(n)
	fmt.Println(result)
}

func viralAdverts(n int) int {
	people := 5
	likes := 2
	totalLikes := 2
	for i := 1; i < n; i++ {
		people = likes * 3
		likes = people / 2
		totalLikes += likes
	}
	return totalLikes
}
