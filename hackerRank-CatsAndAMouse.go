package main

import (
	"fmt"
)

func main() {
	var q, x, y, z int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&x, &y, &z)
		result := catsAndAMouse(x, y, z)
		fmt.Println(result)
	}
}

func catsAndAMouse(x, y, z int) string {
	xDiff := abs(x-z)
	yDiff := abs(y-z)
	if xDiff < yDiff {
		return "Cat A"
	} else if yDiff < xDiff {
		return "Cat B"
	} else { // equal
		return "Mouse C"
	}
}

func abs(x int) int {
	if x < 0 {
		return -1*x
	}
	return x
}
