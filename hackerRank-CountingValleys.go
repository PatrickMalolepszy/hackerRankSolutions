package main

import (
	"fmt"
)

func main() {
	var n int
	var steps string
	fmt.Scan(&n, &steps)
	result := valleyCount(steps)
	fmt.Println(result)
}

func valleyCount(steps string) int {
	seaLevel := 0
	valleyCount := 0
	canBeValley := false
	for i := range steps {
		if steps[i] == 'U' {
			seaLevel++
			if seaLevel == 0 && canBeValley {
				valleyCount++
			}
		} else { // down
			seaLevel--
			if seaLevel < 0 {
				canBeValley = true
			}
		}
	}
	return valleyCount
}