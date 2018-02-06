package main

import (
	"fmt"
	"sort"
)

func main() {
	var t, n, a, b int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n, &a, &b)
		result := manasaStones(n, a, b)
		for i := range result {
			fmt.Print(result[i], " ")
		}
		fmt.Println()
	}
}

func manasaStones(n, a, b int) []int {
	resultSet := make(map[int]bool)
	resultSet[0] = true
	for i := 0 ; i < n-1 ; i++ {
		temp := make(map[int]bool)
		for k := range resultSet {
			delete(resultSet, k)
			temp[k+a] = true
			temp[k+b] = true
		}
		resultSet = temp
	}
	resultArray := make([]int, len(resultSet))
	i := 0
	for k := range resultSet {
		resultArray[i] = k
		i++
	}
	sort.Ints(resultArray)
	return resultArray
}
