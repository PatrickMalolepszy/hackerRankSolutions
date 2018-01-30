package main

import "fmt"

func main() {
	var t, b, w, x, y, z int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&b, &w, &x, &y, &z)
		result := taumBday(b, w, x, y, z)
		fmt.Println(result)
	}
}

func taumBday(b, w, x, y, z int) int {
	if x + z < y { //cheaper to buy and convert black->white
		return b*x + w*(x+z)
	} else if y	 + z < x { //cheaper to buy and convert white->black
		return w*y + b*(y+z)
	} else { //just buy the gifts
		return b*x + w*y
	}
}
