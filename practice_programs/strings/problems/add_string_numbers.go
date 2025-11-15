package problems

import (
	"fmt"
	"strconv"
)

/*
Time complexity -> O(n)
Space -> O(n)
*/
func AddStringNumbers(num1 string, num2 string) string {
	carry := 0
	sum := ""
	i := len(num1) - 1
	j := len(num2) - 1

	for i >= 0 || j >= 0 || carry > 0 {
		digit1 := 0
		if i >= 0 {
			digit1 = int(num1[i] - '0')
		}

		digit2 := 0
		if j >= 0 {
			digit2 = int(num2[j] - '0')
		}

		temp := digit1 + digit2 + carry
		sum += strconv.Itoa(temp % 10)
		carry = temp / 10

		i--
		j--
	}

	runes := []rune(sum)

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}

	return string(runes)
}

func RunAddStringNumbersTest() {
	num1 := "11"
	num2 := "123"

	response := AddStringNumbers(num1, num2)
	fmt.Printf("the sum is %s", response)
}
