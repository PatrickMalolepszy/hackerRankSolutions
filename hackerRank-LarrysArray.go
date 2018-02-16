package main

import (
	"fmt"
	"sort"
)

func main() {
	var t, n int
	fmt.Scan(&t)
	for i := 0 ; i < t ;i++ {
		fmt.Scan(&n)
		a := make([]int, n)
		for j := 0 ; j < n ; j++ {
			fmt.Scan(&a[j])
		}
		if larryArray(n, a) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func rotate(i, j int, a []int)  {
	t := a[i]
	a[i] = a[j]
	a[j] = a[i+1]
	a[i+1] = t
}

func larryArray(n int, a []int) bool {
	done := false
	for !done {
		done = true
		for i := 0 ; i < n - 2 ; i++ {
			for (a[i] > a[i+1]) || (a[i] > a[i+2]) {
				done = false
				rotate(i, i+2, a)
			}
		}
	}
	b := make([]int, n)
	copy(b, a)
	sort.Ints(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}