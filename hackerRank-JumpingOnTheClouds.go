package main

import "fmt"

func main() {
	var n, c int
	fmt.Scan(&n)
	clouds := make([]bool, n)
	for i := range clouds {
		fmt.Scan(&c)
		if c == 1 {
			clouds[i] = true
		} else {
			clouds[i] = false
		}
	}
	result := jumpingOnTheClouds(clouds)
	fmt.Println(result)
}

func jumpingOnTheClouds(clouds []bool) int {
	jumpCount := 0
	curPos := 0
	for curPos < len(clouds) - 2 {
		jumpCount++
		if !clouds[curPos+2] {
			curPos += 2
		} else {
			curPos++
		}
	}
	if curPos == len(clouds) - 2 {
		jumpCount++
	}
	return jumpCount
}
