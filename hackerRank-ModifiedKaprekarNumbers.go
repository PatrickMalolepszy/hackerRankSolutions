package main

import (
	"fmt"
	"strconv"
)

func main() {
	var p, q int
	fmt.Scan(&p, &q)
	invalid := true
	for i := p ; i <= q ; i++ {
		if isKaprekar(i) {
			invalid = false
			fmt.Print(i)
			if i != q {
				fmt.Print(" ")
			}
		}
	}
	if invalid {
		fmt.Println("INVALID RANGE")
	}
}

func isKaprekar(n int) bool {
	n2 := n*n
	sn2 := strconv.Itoa(n2)
	left := sn2[:len(sn2)/2]
	right := sn2[len(sn2)/2:]
	leftNum, _ := strconv.Atoi(left)
	rightNum, _ := strconv.Atoi(right)
	if leftNum + rightNum == n {
		return true
	}
	return false
}
