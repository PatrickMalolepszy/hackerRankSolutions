package main

import "fmt"

func main() {
	charHeight := make(map[byte]int)
	var h int
	var in string
	for i := 0; i < 26; i++ {
		fmt.Scan(&h)
		charHeight[byte('a'+i)] = h
	}
	fmt.Scan(&in)
	maxHeight := 0
	for i := range in {
		if charHeight[in[i]] > maxHeight {
			maxHeight = charHeight[in[i]]
		}
	}
	area := maxHeight * len(in)
	fmt.Println(area)
}
