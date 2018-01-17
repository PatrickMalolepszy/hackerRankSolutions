package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n, &p)
	result := drawingBook(n, p)
	fmt.Println(result)
}

// Find minimum number of page turns to get to page p
// Book has n pages, can start from back or front of book.
func drawingBook(n, p int) int {
	diff := p
	startCount := diff / 2

	diff = n - p
	if n % 2 == 0 {
		diff++
	}
	endCount := diff / 2

	if startCount < endCount {
		return startCount
	} else {
		return endCount
	}
}