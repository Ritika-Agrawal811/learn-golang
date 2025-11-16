package problems

import (
	"fmt"
	"strconv"
)

/*
date format is YYYY-MM-DD
  - go allows index slicing for strings - end index is excluded

Time and Space complexity are both O(1).
*/
func DayOfTheYear(date string) int {
	period := []int{
		0,
		31,
		59,
		90,
		120,
		151,
		181,
		212,
		243,
		273,
		304,
		334,
	}

	day, _ := strconv.Atoi(date[8:10])
	month, _ := strconv.Atoi(date[5:7])
	year, _ := strconv.Atoi(date[0:4])

	daysCount := day + period[month-1]

	if isEvenYear(year) && month > 2 {
		daysCount += 1
	}

	return daysCount

}

/*
A year is leap if:
  - Divisible by 400 → always leap
  - Divisible by 4 but not divisible by 100 → leap
  - Otherwise → not leap
*/
func isEvenYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func RunDayOfTheYearTest() {
	date := "2000-12-04"
	response := DayOfTheYear(date)
	fmt.Printf("day number is: %d", response)
}
