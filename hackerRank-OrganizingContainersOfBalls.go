package main

import "fmt"

func main() {
	var q, n int
	fmt.Scan(&q)
	for i := 0 ; i < q ; i++ {
		fmt.Scan(&n)
		matrix := make([][]int, n)
		for j := range matrix {
			matrix[j] = make([]int, n)
		}
		for j := 0 ; j < n; j++ {
			for k := 0 ; k < n; k++ {
				fmt.Scan(&matrix[j][k])
			}
		}
		result := organizingContainers(matrix)
		if result {
			fmt.Println("Possible")
		} else {
			fmt.Println("Impossible")
		}
	}
}

// not about min/optimization - just can be done or not.
func organizingContainers(m [][]int) bool {
	sumsSameColour := make([]int, len(m))
	containerSum := make([]int, len(m))
	for i := 0 ; i < len(m) ; i++ {
		for j := 0 ; j < len(m) ; j++ {
			sumsSameColour[i] += m[j][i]
			containerSum[i] += m[i][j]
		}
	}

	found := true
	for i := 0 ; i < len(m) && found ; i++ {
		found = false
		for j := 0 ; j < len(m) ; j++ {
			if containerSum[j] == sumsSameColour[i] {
				found = true
				containerSum[j] = -1
				sumsSameColour[i] = -1
			}
		}
		if !found {
			return false
		}
	}
	return true
}
