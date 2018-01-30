package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	var n, k, q1, q2, ox, oy int
	fmt.Scan(&n, &k, &q1, &q2)
	q1--
	q2--
	scanner := bufio.NewScanner(os.Stdin)
	obst := make([][]bool, n)
	for i := range obst {
		obst[i] = make([]bool, n)
	}
	for i := 0 ; i < k ; i++ {
		scanner.Scan()
		//fmt.Println(scanner.Text())
		split := strings.Split(scanner.Text(), " ")
		if len(split) != 2 {
			continue
		}
		ox, _ = strconv.Atoi(split[0])
		oy, _ = strconv.Atoi(split[1])
		//fmt.Println("obst at x:", ox-1, " y:", oy-1)
		obst[ox-1][oy-1] = true
	}
	result := queenAttack(n, q1, q2, obst)
	fmt.Println(result)
}

func queenAttack(n, qX, qY int, obst [][]bool) int {
	count := 0
	// left:
	for i := qX - 1 ; i >= 0 ; i-- {
		if !obst[i][qY] {
			count++
		} else {
			break
		}
	}
	// right:
	for i := qX + 1 ; i < n ; i++ {
		if !obst[i][qY] {
			count++
		} else {
			break
		}
	}
	// down:
	for i := qY - 1 ; i >= 0 ; i-- {
		if !obst[qX][i] {
			count++
		} else {
			break
		}
	}
	// up:
	for i := qY + 1 ; i < n ; i++ {
		if !obst[qX][i] {
			count++
		} else {
			break
		}
	}
	// up & right:
	for i := 1 ; qY - i >= 0 && qX + i < n && !obst[qX+i][qY-i] ; i++ {
		count++
	}
	// up & left:
	for i := 1 ; qY - i >= 0 && qX - i >= 0 && !obst[qX-i][qY-i] ; i++ {
		count++
	}
	// down & right:
	for i := 1 ; qY + i < n && qX + i < n && !obst[qX+i][qY+i] ; i++ {
		count++
	}
	// down & left:
	for i := 1 ; qY + i < n && qX - i >= 0 && !obst[qX-i][qY+i] ; i++ {
		count++
	}
	return count
}
