package main

import "fmt"

func main() {
	var t, n, k int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n, &k)
		result, err := absolutePerm(n, k)
		if err != nil {
			fmt.Print(-1)
		} else {
			for j := range result {
				fmt.Print(result[j], " ")
			}
		}
		fmt.Println()
	}
}

func absolutePerm(n, k int) ([]int, error) {
	if k > n/2 {
		return nil, fmt.Errorf("k is larger then n/2")
	}
	result := make([]int, n)
	for i := range result {
		result[i] = i + 1
	}
	if k == 0 {
		return result, nil
	}
	if (n%(k*2) != 0) || (n%2 != 0) {
		return []int{-1}, nil
	}
	for i := 0; i < n-k; i += k * 2 {
		for j := i; j < k+i; j++ {
			result[j], result[k+j] = result[k+j], result[j]
		}
	}
	return result, nil
}
