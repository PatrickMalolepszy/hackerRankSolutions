package main

import (
	"fmt"
)

var hours = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
}

func main() {
	var h, m int
	fmt.Scan(&h, &m)
	numWord := hours
	result := timeInWords(h, m, numWord)
	fmt.Println(result)
}

func timeInWords(h, m int, numWord map[int]string) string {
	var result string

	if m < 30 {
		if m <= 20 {
			result = fmt.Sprintf("%s minutes past %s", numWord[m], numWord[h])
		} else {
			result = fmt.Sprintf("twenty %s minutes past %s", numWord[m%10], numWord[h])
		}
	} else {
		val := 60 - m
		if val <= 20 {
			result = fmt.Sprintf("%s minutes to %s", numWord[val], numWord[(h+1)%12])
		} else {
			result = fmt.Sprintf("twenty %s minutes to %s", numWord[val%10], numWord[(h+1)%12])
		}
	}

	switch m {
	case 0:
		result = fmt.Sprintf("%s o' clock", numWord[h])
	case 1:
		result = fmt.Sprintf("%s minute past %s", numWord[m], numWord[h])
	case 15:
		result = fmt.Sprintf("quarter past %s", numWord[h])
	case 30:
		result = fmt.Sprintf("half past %s", numWord[h])
	case 45:
		result = fmt.Sprintf("quarter to %s", numWord[(h+1)%12])
	case 59:
		result = fmt.Sprintf("%s minute to %s", numWord[m], numWord[(h+1)%12])
	}

	return result
}
//
//func getNumberWords() map[int]string {
//	input, err := os.Open("numData.txt")
//	scanner := bufio.NewScanner(input)
//	numWords := make(map[int]string)
//	if err != nil {
//		fmt.Println("could not open file: ", err)
//		return nil
//	}
//	for scanner.Scan() {
//		parts := strings.Split(scanner.Text(), " ")
//		num, err := strconv.Atoi(parts[0])
//		if err != nil {
//			fmt.Println("could not atoi: ", err)
//		}
//		numWords[num] = parts[1]
//	}
//	return numWords
//}
