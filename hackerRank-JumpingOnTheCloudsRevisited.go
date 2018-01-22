package main

import "fmt"

func main() {
	var n, k, c int
	fmt.Scan(&n, &k)
	clouds := make([]bool, n)
	for i := range clouds {
		fmt.Scan(&c)
		if c == 1 {
			clouds[i] = true
		} else {
			clouds[i] = false
		}
	}
	result := cloudJump(k, clouds)
	fmt.Println(result)
}

func cloudJump(k int, clouds []bool) int {
	energy := 100
	currentCloud := 0
	for true {
		energy--
		currentCloud = (currentCloud + k) % len(clouds)
		if clouds[currentCloud] {
			energy -= 2
		}
		if currentCloud == 0 {
			return energy
		}
	}
	return -1
}