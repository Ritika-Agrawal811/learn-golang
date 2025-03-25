package main

import "fmt"

/* should be outside of main() as main() is also a function */
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

/* named return values */
func rectangleArea(length, width int) (area int) {
	area = length * width // `area` is already declared
	return                // No need to specify `area` explicitly
}

func main() {
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	fmt.Println("Area:", rectangleArea(5, 10))

	/* ignoring second returned value */
	q1, _ := divide(18, 3)
	fmt.Println("Quotient:", q1)
}
