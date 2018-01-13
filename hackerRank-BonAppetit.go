package main

import "fmt"

func main() {
	var n, k, bill int
	fmt.Scan(&n, &k)
	food := make([]int, n)
	for i := 0 ; i < n ; i++ {
		fmt.Scan(&food[i])
	}
	fmt.Scan(&bill)
	result := bonAppétit(k, bill, food)
	if result == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(result)
	}

}

func bonAppétit(k, bill int, food []int) int {
	sum := 0
	for i := range food {
		sum += food[i]
	}
	sum -= food[k]
	herShare := sum/2
	diff := bill - herShare
	if diff <= 0 {
		return 0
	} else {
		return diff
	}
}