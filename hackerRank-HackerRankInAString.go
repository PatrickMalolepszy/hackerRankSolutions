package main

import "fmt"

func main() {
	var n int
	var t string
	fmt.Scan(&n)
	for i := 0 ; i < n ; i++ {
		fmt.Scan(&t)
		if hackerRankInAString(t) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func hackerRankInAString(s string) bool {
	hackerRank := "hackerrank"
	pos := 0
	for i := range s {
		if pos == len(hackerRank) {
			return true
		}
		if s[i] == hackerRank[pos] {
			pos++
		}
	}
	return pos == len(hackerRank)
}