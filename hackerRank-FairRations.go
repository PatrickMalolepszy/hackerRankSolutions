package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	bread := make([]int, n)
	for i := range bread {
		fmt.Scan(&bread[i])
	}
	result := rankHackerCastle(bread)
	if result == -1 {
		fmt.Println("NO")
	} else {
		fmt.Println(result)
	}
}

func rankHackerCastle(bread []int) int {
	count := 0
	// first
	if bread[0] % 2 != 0 {
		count++
		bread[0]++
		bread[1]++
	}
	for i := 1 ; i < len(bread) - 1 ; i++ {
		if bread[i] % 2 != 0 {
			count++
			bread[i]++
			if bread[i-1] % 2 != 0 {
				bread[i-1]++
			} else {
				bread[i+1]++
			}
		}
	}
	// check last
	if bread[len(bread)-1] % 2 != 0 {
		return -1
	}
	return count*2
}