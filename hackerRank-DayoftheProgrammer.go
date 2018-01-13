package main

import "fmt"

func main() {
	var y int
	fmt.Scan(&y)
	result := DayOfTheProgrammer(y)
	fmt.Println(result)
}

func DayOfTheProgrammer(year int) string {
	leapYearDate := fmt.Sprintf("12.09.%d", year)
	regularYearDate := fmt.Sprintf("13.09.%d", year)
	if year <= 1917 { // julian cal.
		if year % 4 == 0 {
			return leapYearDate
		} else {
			return regularYearDate
		}
	} else if year >= 1919 { // gregorian cal.
		if year % 400 == 0 || (year % 4 == 0 && year % 100 != 0 ) {
			return leapYearDate
		} else {
			return regularYearDate
		}
	} else { // special case: year = 1918
		return "26.09.1918"
	}
}