package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	numStr := strconv.Itoa(x)

	for i := 0; i <= len(numStr)/2; i++ {
		checkIndex := len(numStr) - 1 - i

		if numStr[i] != numStr[checkIndex] {
			return false
		}
	}

	return true
}

// this function checks for palindrome without converting the number to string */
func checkPalindromeWithoutConvertingToString(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return x == reversed || x == reversed/10
}

func main() {

	testCases := []int{121, -121, 10, 444, 5034, 56765}

	for _, num := range testCases {
		result := isPalindrome(num)
		fmt.Println(result)

		response := checkPalindromeWithoutConvertingToString(num)
		fmt.Printf("checking without converting to string: %v\n", response)

	}

}
