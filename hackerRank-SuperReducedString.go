package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	result := superReducedString(s)
	if len(result) == 0 {
		fmt.Println("Empty String")
	} else {
		fmt.Println(result)
	}
}

func superReducedString(s string) string {
	done := false
	result := ""
	for !done && len(s) > 0 {
		prevCount := 1
		done = true
		for i := 1; i < len(s); i++ {
			prev, cur := s[i-1], s[i]
			if prev == cur {
				prevCount++
				done = false
			} else {
				if prevCount%2 != 0 {
					result += string(prev)
				}
				prevCount = 1
			}
		}
		if prevCount%2 != 0 {
			result += string(s[len(s)-1])
		}
		s = result
		result = ""
	}
	return s
}
