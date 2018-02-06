package main

import (
	"fmt"
)

func main() {
	var g, n int
	fmt.Scan(&g)
	var game string
	for i := 0 ; i < g ; i++ {
		fmt.Scan(&n, &game)
		if happyLadybugs(n, game) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func happyLadybugs(n int, game string) bool {
	gameState := make(map[byte]int)
	for i := 0 ; i < n ; i++ {
		gameState[game[i]]++
	}
	space := false
	for k, v := range gameState {
		if k == '_' {
			space = true
			continue
		}
		if v == 1 {
			return false
		}
	}
	if space {
		return true
	} else {
		for i := 1 ; i < n-1 ; i++ {
			if game[i-1] != game[i] && game[i] != game[i+1] {
				return false
			}
		}
	}
	return true
}