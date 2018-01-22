package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var n int
	fmt.Scan(&n)
	result := repeatedString(s, n)
	fmt.Println(result)
}

func repeatedString(s string, n int) int {
	result := 0
	aCountInWord := 0
	for i := range s {
		if s[i] == 'a' {
			aCountInWord++
		}
	}
	fullWordCount := n / len(s)
	result += fullWordCount * aCountInWord
	leftOver := n % len(s)
	for i := 0 ; i < leftOver ; i++ {
		if s[i] == 'a' {
			result++
		}
	}
	return result

}