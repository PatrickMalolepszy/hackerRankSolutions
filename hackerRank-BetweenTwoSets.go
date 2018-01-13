package main

import "fmt"

// just getting things ready.

func main() {

	var n,m int
	fmt.Scan(&n, &m)

	a := make([]int, n)
	b := make([]int, m)

	for i := 0 ; i < n ; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0 ; i < m ; i++ {
		fmt.Scan(&b[i])
	}

	count := 0
	for i := 1 ; i <= 100 ; i++ {
		flag := true
		for j := range a {
			if i % a[j] != 0 {
				flag = false
				break
			}
		}
		if flag {
			for j := range b {
				if b[j] % i != 0 {
					flag = false
					break
				}
			}
		}
		if flag {
			count++
		}
	}
	fmt.Println(count)
}