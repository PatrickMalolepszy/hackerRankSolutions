package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	result := strangeCounter(t)
	fmt.Println(result)
}

func strangeCounter(t int) int {
	curCountDown := 3
	curTime := 0
	for {
		if t <= curTime + curCountDown {
			total := curTime + curCountDown
			return total - t + 1
		}
		curTime += curCountDown
		curCountDown *= 2
	}
}