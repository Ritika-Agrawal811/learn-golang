package numbers

import "fmt"

/* An Armstrong number is a number that is equal to the
sum of cubes of its digits. Example: 153 = 1³ + 5³ + 3³ */

func CheckArmstrongNumber(num int) bool {
	sum := 0

	copy := num
	for copy > 0 {
		digit := copy % 10
		sum += digit * digit * digit
		copy /= 10
	}

	return sum == num
}

func RunArmstrongNumberTest() {
	cases := []int{153, 444}

	for _, num := range cases {
		result := CheckArmstrongNumber(num)
		fmt.Printf("%d is armstrong: %v\n", num, result)
	}
}
