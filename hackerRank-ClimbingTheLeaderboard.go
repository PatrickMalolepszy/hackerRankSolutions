package main

import "fmt"

func main() {
	var n,m int
	fmt.Scan(&n)
	scores := make([]int,n)
	for i := range scores {
		fmt.Scan(&scores[i])
	}
	fmt.Scan(&m)
	alice := make([]int, m)
	for i := range alice {
		fmt.Scan(&alice[i])
	}
	result := climbingTheLeaderboard(scores, alice)
	for i := range result {
		fmt.Println(result[i])
	}
}

func climbingTheLeaderboard(scores, alice []int) []int {
	result := make([]int, len(alice))

	scoreSet := make(map[int]bool)
	for i := range scores {
		scoreSet[scores[i]] = true
	}
	uniqueScores := len(scoreSet)

	cur := len(scores) - 1
	currentPlace := uniqueScores + 1

	for i := range alice {
		for cur >= 0 && alice[i] > scores[cur] {
			if cur == len(scores) - 1 {
				currentPlace--
			} else if scores[cur+1] != scores[cur] {
				currentPlace--
			}
			cur--
		}
		if cur < 0 { // first place
			result[i] = 1
		} else if alice[i] == scores[cur] {
			result[i] = currentPlace - 1
		} else { // less then current score.
			result[i] = currentPlace
		}
	}
	return result
}
