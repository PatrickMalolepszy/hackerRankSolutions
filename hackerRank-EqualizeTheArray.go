package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	result := equalizeTheArray(arr)
	fmt.Println(result)
}

func equalizeTheArray(a []int) int {
	numCounts := make(map[int]int)
	for i := range a {
		numCounts[a[i]]++
	}
	max := 0
	for _, v := range numCounts {
		if max < v {
			max = v
		}
	}
	return len(a) - max
}