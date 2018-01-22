package main

import "fmt"

func main() {
	var n, k, q1, q2, ox, oy int
	fmt.Scan(&n, &k, &q1, &q2)
	obst := make([][]bool, n)
	for i := range obst {
		obst[i] = make([]bool, n)
	}
	for i := 0 ; i < k ; i++ {
		fmt.Scan(&ox, &oy)
		obst[ox-1][oy-1] = true
	}
	result := queenAttack(n, q1, q2, obst)
	fmt.Println(result)
}

func queenAttack(n, qX, qY int, obst [][]bool) int {
	count := 0
	qX -= 1
	qY -= 1
	// left:
	for i := qX - 1 ; i >= 0 && !obst[i][qY] ; i-- {
		count++
	}
	// right:
	for i := qX + 1 ; i < n && !obst[i][qY] ; i++ {
		count++
	}
	// up:
	for i := qY - 1 ; i >= 0 && !obst[qX][i] ; i-- {
		count++
	}
	// down:
	for i := qY + 1 ; i < n && !obst[qX][i] ; i++ {
		count++
	}
	// up & right:
	for i := 1 ; qY - i >= 0 && qX + i < n && !obst[qX+i][qY-1] ; i++ {
		count++
	}
	// up & left:
	for i := 1 ; qY - i >= 0 && qX - i >= 0 && !obst[qX-i][qY-1] ; i++ {
		count++
	}
	// down & right:
	for i := 1 ; qY + i < n && qX + i < n && !obst[qX+i][qY+1] ; i++ {
		count++
	}
	// down & left:
	for i := 1 ; qY + i < n && qX - i >= 0 && !obst[qX-i][qY+1] ; i++ {
		count++
	}
	return count
}
