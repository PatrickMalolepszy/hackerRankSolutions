package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	array := make([]int, n)
	for i := range array {
		fmt.Scan(&array[i])
	}
	result := pickingNumbers(array)
	fmt.Println(result)
}

func pickingNumbers(a []int) int {
	numberCounts := make(map[int]int)
	for i := range a {
		numberCounts[a[i]]++
	}
	max := 0
	for i := 0 ; i < 100 ; i++ {
		count := numberCounts[i-1] + numberCounts[i]
		if count > max {
			max = count
		}
		count = numberCounts[i+1] + numberCounts[i]
		if count > max {
			max = count
		}
	}
	return max
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
