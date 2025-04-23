package numbers

import "fmt"

/* Time Complexity is ---> O(n) */
func Factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}

	factorial := 1
	for num >= 2 {
		factorial *= num
		num--
	}

	return factorial
}

func RunFactorialTest() {
	cases := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range cases {
		fact := Factorial(num)
		fmt.Printf("factorial of %d is %d \n", num, fact)
	}
}
