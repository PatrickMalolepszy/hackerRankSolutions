package main

import (
	. "math/big"
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	result := extraLongFactorials(n)
	fmt.Println(result)
}

func extraLongFactorials(n int) *Int {
	x := NewInt(int64(n))
	for i := n-1 ; i >= 2 ; i-- {
		x = x.Mul(x, NewInt(int64(i)))
	}
	return x
}