package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var in string
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&in)
		if balancedBrackets(in) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

// bracketSet : {}<>()
func balancedBrackets(brackets string) bool {
	//var opening, closing map[int32]bool
	bracketSet := "{}()[]<>"
	bracketStack := NewStack()
	for i := range brackets {
		pos := strings.Index(bracketSet, string(brackets[i]))
		if pos == -1 {
			return false // bracket not defined
		}
		if bracketSet[pos] % 2 == 0 { // opening bracket
			bracketStack.Push(brackets[i])
		} else { // closing bracket
			err, top := bracketStack.Pop()
			if err != nil || top != bracketSet[pos-1] {
				return false
			}
			break
		}
	}
	return bracketStack.IsEmpty()
}
