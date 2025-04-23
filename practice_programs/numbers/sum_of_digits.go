package numbers

import "fmt"

/* Time Complexity is ---> O(d) where d is the number of digits */
func SumOfDigits(num int) int {

	if num == 0 {
		return 0
	}

	/* handle negative numbers */
	if num < 0 {
		num = -num
	}

	sum := 0

	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum
}

func RunSumOfDigitsTest() {
	cases := []int{122, 345, 333, 890012}

	for _, num := range cases {
		result := SumOfDigits(num)
		fmt.Printf("sum of digits for %d is %d \n", num, result)
	}
}
