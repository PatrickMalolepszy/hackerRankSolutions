package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	var t int
	var w string
	fmt.Scan(&t)
	for i:=0 ; i < t ; i++ {
		scanner.Scan()
		w = scanner.Text()
		result := biggerIsGreater(w)
		output.WriteString(result + "\n")
	}
	output.Flush()
}

func biggerIsGreater(w string) string {
	err, result := nextLexicographicalPermutation(w)
	if err != nil {
		return "no answer"
	}
	return result
}

func nextLexicographicalPermutation(s string) (error, string) {
	if len(s) <= 1 {
		return fmt.Errorf("no next permutation"), ""
	}
	w := []byte(s)
	// find pivot:
	pivot := -1
	for i := len(w) - 1 ; i > 0 ; i-- {
		if w[i-1] < w[i] {
			pivot = i-1
			break
		}
	}

	if pivot == -1 {
		return fmt.Errorf("no next permutation"), ""
	}

	// find rightmost element greater than pivot:
	for i := len(w) - 1 ; i > 0 ; i-- {
		if w[pivot] < w[i] {
			w[pivot], w[i] = w[i], w[pivot]
			break
		}
	}

	// reverse suffix:
	for i, j := pivot+1, len(w)-1 ; i < j ; i,j = i+1,j-1 {
		w[i], w[j] = w[j], w[i]
	}

	return nil, string(w)
}

