package main

import "fmt"

func main() {
	var n, d int
	fmt.Scan(&n, &d)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	result := beautifulTriplets(d, a)
	fmt.Println(result)
}

func beautifulTriplets(d int, a []int) int {
	count := 0
	middleDigits := make(map[int]int)
	for i := 0 ; i < len(a) ; i++ {
		for j := i+1 ; j < len(a) ; j++ {
			if a[j] - a[i] == d {
				count += middleDigits[i]
				middleDigits[j]++
			}
		}
	}
	return count
}
