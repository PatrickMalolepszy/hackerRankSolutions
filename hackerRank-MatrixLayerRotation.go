package main

import "fmt"

func main() {
	var m, n, r int
	fmt.Scan(&m, &n, &r)
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Scan(&matrix[i][j])
		}
	}
	result := layerRotation(m, n, matrix, r)
	for i := range result {
		for j := range result[i] {
			fmt.Print(result[i][j], " ")
		}
		fmt.Println()
	}
}

func layerRotation(m, n int, matrix [][]int, r int) [][]int {
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := range matrix {
		for j := range matrix[i] {
			result[i][j] = matrix[i][j]
		}
	}


	minX, minY, maxX, maxY := 0, 0, m, n
	for (minX < maxX) && (minY < maxY) {

		fullCircle := 2*(maxX-minX-1) + 2*(maxY-minY-1)
		rotations := r % fullCircle
		//fmt.Println(rotations, fullCircle)

		for i := 0 ; i < rotations ; i++ {
			rotateMatrix2(minX, maxX, minY, maxY, result)
		}

		minX++
		minY++
		maxX--
		maxY--
	}


	return result
}

func rotateMatrix2(minX, maxX, minY, maxY int, matrix [][]int) {
	t := matrix[minX][minY]
	// top row
	for i := minX; i < maxY-1; i++ {
		matrix[minX][i] = matrix[minX][i+1]
	}
	// left col
	t2 := matrix[maxX-1][minY]
	for i := maxX - 2; i >= minX + 1; i-- {
		matrix[i+1][minY] = matrix[i][minY]
	}
	matrix[minX + 1][minY] = t
	// bottom row
	t3 := matrix[maxX-1][maxY-1]
	for i := maxY - 1; i >= minY + 1; i-- {
		matrix[maxX-1][i] = matrix[maxX-1][i-1]
	}
	matrix[maxX-1][minY+1] = t2
	// right col
	for i := minX; i < maxX-2; i++ {
		matrix[i][maxY-1] = matrix[i+1][maxY-1]
	}
	matrix[maxX-2][maxY-1] = t3
}
