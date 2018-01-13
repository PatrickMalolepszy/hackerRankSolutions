/*
https://www.hackerrank.com/challenges/birthday-cake-candles/problem
 */
package main

import "fmt"

func main() {
	var i, j int
	fmt.Scan(&i)
	max := 1
	candles := make([]int, i)
	for n := 0; n < i; n++ {
		fmt.Scan(&j)
		if j > max {
			max = j
		}
		candles[n] = j
	}
	count := 0
	for n := 0; n < i; n++ {
		if candles[n] == max {
			count++
		}
	}
	fmt.Println(count)
}
