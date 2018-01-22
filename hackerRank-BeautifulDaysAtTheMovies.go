package main

import (
	"fmt"
	"unicode/utf8"
	"strconv"
)

func main() {
	var i, j, k int
	fmt.Scan(&i, &j, &k)
	result := beautifulDays(i, j, k)
	fmt.Println(result)
}

// count the number of days where the reverse num diff = k
func beautifulDays(i, j, k int) int {
	count := 0
	for s := i ; s <= j ; s++ {
		//fmt.Println(s, reverse(s))
		if (s - reverse(s)) % k == 0{
			count++
		}
	}
	return count
}

func reverse(x int) int {
	var s string
	s = strconv.Itoa(x)
	s = ReverseString(s)
	result, _ := strconv.Atoi(s)
	return result
}

func ReverseString(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
