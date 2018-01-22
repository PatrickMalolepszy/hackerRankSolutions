package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	seq := make([]int, n)
	for i := range seq {
		fmt.Scan(&seq[i])
	}
	result := seqEq(seq)
	for i := range result {
		fmt.Println(result[i])
	}
}

func seqEq(seq []int) []int {
	result := make([]int, len(seq))
	for i := range seq {
		result[i] = Px(Px(i + 1, seq), seq)
	}
	return result
}

func Px(x int, seq []int) int {
	for i := range seq {
		if seq[i] == x {
			return i + 1
		}
	}
	return -1
}