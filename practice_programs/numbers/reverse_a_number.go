package numbers

import "fmt"

/* Time Complexity is ---> O(d) where d is the number of digits */
func ReverseNumber(num int) int {

	isNegative := false

	if num < 0 {
		isNegative = true
		num = -num
	}

	reversed := 0

	for num > 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}

	if isNegative {
		return -reversed
	}

	return reversed
}

func RunReversedNumberTest() {
	cases := []int{12, 345, 647, 29850}

	for _, num := range cases {
		sum := ReverseNumber(num)
		fmt.Printf("number : %d, reversed : %d \n", num, sum)
	}
}
