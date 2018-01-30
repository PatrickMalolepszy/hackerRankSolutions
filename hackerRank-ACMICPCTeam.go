package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	topics := make([]string, n)
	for i := range topics {
		fmt.Scan(&topics[i])
	}
	maxTopics, numTeams := acmTeam(topics)
	fmt.Println(maxTopics)
	fmt.Println(numTeams)
}

// O(n^2*m)
func acmTeam(topics []string) (int,int) {
	maxTopics := 0
	topicCounts := make(map[int]int)
	for i := 0 ; i < len(topics) ; i++ {
		for j := i+1 ; j < len(topics) ; j++ {
			topicSet := make(map[int]bool)
			for k := range topics[i] {
				if topics[i][k] == '1' || topics[j][k] == '1' {
					topicSet[k] = true
				}
			}
			curTeamTopics := len(topicSet)
			topicCounts[curTeamTopics]++
			if maxTopics < curTeamTopics {
				maxTopics = curTeamTopics
			}
		}
	}
	return maxTopics, topicCounts[maxTopics]
}
