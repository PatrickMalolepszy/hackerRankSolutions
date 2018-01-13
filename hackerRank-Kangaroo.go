package main

import "fmt"

func main() {
	var x1, v1, x2, v2 int
	fmt.Scan(&x1, &v1, &x2, &v2)
	if v2 >= v1 {
		fmt.Println("NO")
		return
	}
	for x1 < x2 {
		x1 += v1
		x2 += v2
	}
	if x1 == x2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
