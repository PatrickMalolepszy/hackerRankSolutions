package main

import "fmt"

func main() {
	var s, t string
	var k int
	fmt.Scan(&s, &t, &k)
	result := appendAndDelete(s, t, k)
	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func appendAndDelete(s, t string, k int) bool {
	if deleteWhole(s, t, k) {
		//fmt.Println("deleted")
		return true
	}
	if subString(s, t, k) {
		//fmt.Println("used substring")
		return true
	}
	return false
}

func deleteWhole(s, t string, k int) bool {
	count := 0
	count += len(s) + len(t)
	if count <= k {
		return true
	}
	return false
}

//hackerhappy
//hackerrank
func subString(s, t string, k int) bool {
	var i int
	for i = 0; i < len(t) && i < len(s); i++ {
		if s[i] != t[i] {
			break
		}
	}
	//fmt.Println(i)
	deletes := len(s) - i
	appends := len(t) - i
	ops := deletes + appends
	if ops == k || ((k-ops) % 2 == 0 && k-ops > 0) {
		return true
	}
	return false
}
