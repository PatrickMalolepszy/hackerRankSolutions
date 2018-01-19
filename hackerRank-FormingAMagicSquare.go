package main

import "fmt"

func main() {
	//var square1 = [][]int{{2, 7, 6},
	//	{9, 5, 1},
	//	{4, 3, 8}}
	//fmt.Println(square1)
	//rotateMatrix(3, square1)
	//fmt.Println(square1)

	magicSquare := make([][]int, 3)
	for i := 0; i < 3; i++ {
		magicSquare[i] = make([]int, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&magicSquare[i][j])
		}
	}
	result := formMagicSquare(magicSquare)
	fmt.Println(result)
}

func formMagicSquare(magicSquare [][]int) int {

	realMagicSquares := make([][][]int, 8)

	realMagicSquares[0] = [][]int{{2, 7, 6},
								{9, 5, 1},
								{4, 3, 8}}

	realMagicSquares[1] = [][]int{{2, 9, 4},
								{7, 5, 3},
								{6, 1, 8}}

	realMagicSquares[2] = [][]int{{4, 9, 2},
								{3, 5, 7},
								{8, 1, 6}}

	realMagicSquares[3] = [][]int{{6, 7, 2},
								{1, 5, 9},
								{8, 3, 4}}

	realMagicSquares[4] = [][]int{{8, 3, 4},
								{1, 5, 9},
								{6, 7, 2}}

	realMagicSquares[5] = [][]int{{8, 1, 6},
								{3, 5, 7},
								{4, 9, 2}}

	realMagicSquares[6] = [][]int{{6, 1, 8},
								{7, 5, 3},
								{2, 9, 4}}

	realMagicSquares[7] = [][]int{{4, 3, 8},
								{9, 5, 1},
								{2, 7, 6}}

	min := 1000
	for i := range realMagicSquares {
		//displaySquare(realMagicSquares[i])
		//fmt.Println()
		cost := 0
		for j := 0 ; j < 3 ; j++ {
			for k := 0 ; k < 3 ; k++ {
				cost += abs(realMagicSquares[i][j][k] - magicSquare[j][k])
			}
		}
		if cost < min {
			min = cost
		}
	}
	return min
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func rotateMatrix(n int, m [][]int) {
	for i := 0 ; i < n /2 ; i++ {
		for j := i ; j < n-i-1 ; j++ {
			t := m[i][j]
			m[i][j] = m[j][n-1-i]
			m[j][n-1-i] = m[n-1-i][n-1-j]
			m[n-1-i][n-1-j] = m[n-1-j][i]
			m[n-1-j][i] = t
		}
	}
}

func displaySquare(m [][]int) {
	for i := range m {
		for j := range m[i] {
			fmt.Print(m[i][j], " ")
		}
		fmt.Println()
	}
}

func copySquare(m [][]int) [][]int {
	newSquare := make([][]int, 3)
	for i := 0 ; i < 3 ; i++ {
		newSquare[i] = make([]int, 3)
	}
	for i := 0 ; i < 3 ; i++ {
		for j := 0 ; j < 3 ; j++ {
			newSquare[i][j] = m[i][j]
		}
	}
	return newSquare
}


