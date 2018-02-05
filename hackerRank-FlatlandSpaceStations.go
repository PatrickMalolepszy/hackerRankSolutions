package main

import (
	"fmt"
)

func main() {
	var n, m, s int
	fmt.Scan(&n, &m)
	spaceStation := make(map[int]bool)
	for i := 0; i < m; i++ {
		fmt.Scan(&s)
		spaceStation[s] = true
	}
	result := flatlandSpaceStations(n, spaceStation)
	fmt.Println(result)
}

func flatlandSpaceStations(n int, spaceStation map[int]bool) int {
	maxD := 0
	leftSpace, rightSpace := -1, -1
	for j := 1; j < n; j++ {
		if spaceStation[j] {
			rightSpace = j
			break
		}
	}
	for i := 0; i < n; i++ {
		if spaceStation[i] {
			leftSpace = i
			found := false
			for j := i+1; j < n; j++ {
				if spaceStation[j] {
					rightSpace = j
					found = true
					break
				}
			}
			if !found {
				rightSpace = -1
			}
			continue
		}
		//r := fmt.Sprintf("At %d with left at %d and right at %d\n", i, leftSpace, rightSpace)
		//fmt.Println(r)
		leftDis, rightDis := 100000, 100000
		if leftSpace != -1 {
			leftDis = i - leftSpace
		}
		if rightSpace != -1 {
			rightDis = rightSpace - i
		}

		if leftDis < rightDis && leftDis > maxD{
			maxD = leftDis
		} else if rightDis > maxD {
			maxD = rightDis
		}
	}
	return maxD
}
