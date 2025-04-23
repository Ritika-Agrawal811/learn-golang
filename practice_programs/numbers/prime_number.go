package numbers

import (
	"fmt"
	"math"
)

/* Time complexity is --> O(n) */
func IsPrime(num int) bool {

	if num <= 1 {
		return false
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

/* Time Complexity is ---> O(√n) */
func CheckPrime(num int) bool {
	if num <= 1 {
		return false
	}

	limit := int(math.Sqrt(float64(num)))

	for i := 2; i <= limit; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func RunPrimeCheckTest() {
	cases := []int{5, 13, 44, 67, 89, 100, -7}

	for _, num := range cases {
		result := CheckPrime(num)
		fmt.Printf("is %d number a prime? ans: %v \n", num, result)
	}
}
