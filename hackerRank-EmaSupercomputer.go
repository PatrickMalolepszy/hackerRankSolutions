package main

import "fmt"

func main() {
	var n, m int
	var line string
	fmt.Scan(&n, &m)
	cells := make([][]bool, n)
	for i := range cells {
		cells[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&line)
		for j := range line {
			if line[j] == 'G' {
				cells[i][j] = true
			}
		}
	}
	result := emaSupercomputer(n, m, cells)
	fmt.Println(result)
}

func emaSupercomputer(n, m int, cells [][]bool) int {
	maxArea := 0
	//count := 0
	for x1 := range cells {
		for y1 := range cells[x1] {
			for x2 := range cells {
				for y2 := range cells[x2] {
					//count++
					//if count%100 == 0 {
					//	//fmt.Println("calculated ", count, " reuslts")
					//}
					// two coor not in same place:
					if x1 == x2 && y1 == y2 {
						continue
					}
					// both coor good cells:
					if !cells[x1][y1] || !cells[x2][y2] {
						continue
					}
					// start expanding:
					i, j := 0,0
					expand1, expand2 := true, true
					for {
						if expand1 {
							expand1Ok := plusExpandOk(x1, y1, n, m, i+1, cells)
							plusHits1 := plusHits(x1, y1, i+1, x2, y2, j)
							//fmt.Println(x1, y1, i, x2, y2, j, expand1Ok, plusHits1)
							if !expand1Ok || plusHits1 {
								// don't expand first plus
								expand1 = false
							} else {
								i++
							}
						}
						if expand2 {
							if !plusExpandOk(x2, y2, n, m, j+1, cells) ||
								plusHits(x2, y2, j+1, x1, y1, i) {
								// don't expand second plus
								expand2 = false
							} else {
								j++
							}
						}
						if !expand1 && !expand2 {
							break
						}
					}
					area := (1 + (i)*4) * (1 + (j)*4)
					//if area == 325 {
					//	fmt.Println("found area:", area, "cur Max:", maxArea)
					//	fmt.Println(x1, y1, x2, y2,"i,j:", i, j)
					//}
					if area > maxArea {
						maxArea = area
					}
					//fmt.Println("max:", maxArea, "i,j: ", i,j)
				}
			}
		}
	}
	return maxArea
}

func plusExpandOk(x, y, n, m, i int, cells [][]bool) bool {
	// right
	if !inBounds(x+i, y, n, m) || !cells[x+i][y] {
		return false
	}
	// left
	if !inBounds(x-i, y, n, m) || !cells[x-i][y] {
		return false
	}
	// up
	if !inBounds(x, y+i, n, m) || !cells[x][y+i] {
		return false
	}
	// down
	if !inBounds(x, y-i, n, m) || !cells[x][y-i] {
		return false
	}
	return true
}

func plusHits(x1, y1, i, x2, y2, j int) bool {
	// right: cell[x1+i][y1]
	if inPlusArea(x1+i, y1, x2, y2, j) {
		return true
	}
	// left
	if inPlusArea(x1-i, y1, x2, y2, j) {
		return true
	}
	// up
	if inPlusArea(x1, y1+i, x2, y2, j) {
		return true
	}
	// down
	if inPlusArea(x1, y1-i, x2, y2, j) {
		return true
	}
	return false
}

// check that single cell x,y is not in plus at x2,y2 expanded by j
func inPlusArea(x, y, x2, y2, j int) bool {
	// horizontal:
	if ((x2 - j) <= x) && (x <= (x2 + j)) && (y == y2) {
		return true
	}
	// vertical:
	if ((y2 - j) <= y) && (y <= (y2 + j)) && (x == x2) {
		return true
	}
	return false
}

func inBounds(x, y, n, m int) bool {
	if (0 <= x) && (x < n) && (0 <= y) && (y < m) {
		return true
	}
	return false
}

// O(n*m)^2