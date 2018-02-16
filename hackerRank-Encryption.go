package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	//fmt.Println(s)
	result := encryption(s)
	fmt.Print(result)
}

func encryption(s string) string {
	L := len(s)
	sqrtL := math.Sqrt(float64(L))
	floorSqrtL := int(sqrtL)
	ceilSqrtL := int(math.Ceil(sqrtL))
	var row, col int
	if floorSqrtL*floorSqrtL >= L {
		row, col = floorSqrtL, floorSqrtL
	} else if floorSqrtL*ceilSqrtL >= L {
		row, col = floorSqrtL, ceilSqrtL
	} else if ceilSqrtL*ceilSqrtL >= L {
		row, col = ceilSqrtL, ceilSqrtL
	} else {
		// todo... err?
	}
	m := make([][]byte, row)
	for i := range m {
		m[i] = make([]byte, col)
	}
	for i := range m {
		for j := range m[i] {
			m[i][j] = 255
		}
	}
	for i := range m {
		for j := range m[i] {
			if (i*col)+j < len(s) {
				m[i][j] = s[(i*col)+j]
			}
		}
	}
	result := ""
	for i := range m[0] {
		for j := range m {
			if m[j][i] == 255 {
				continue
			}
			char := string(m[j][i])
			result += char
		}
		result += " "
	}

	return result[:len(result)-1]
}

func trim(s string) string {
	r := strings.TrimSpace(s)
	r2 := ""
	for i := 1 ; i < len(r) ; i++ {
		if r[i] == ' ' && r[i-1] == ' ' {

		} else {
			r2 += string(r[i-1])
		}
	}
	r2 += string(r[len(r)-1])
	return r2
}
