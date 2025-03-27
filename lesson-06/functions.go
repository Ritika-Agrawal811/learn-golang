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

func checkAccess(user string, hasPermission bool) {
	if user == "" {
		fmt.Println("Invalid user")
		return
	}

	if !hasPermission {
		fmt.Println("Access denied")
		return
	}

	fmt.Println("Access granted")
}

func main() {
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	fmt.Println("Area:", rectangleArea(5, 10))

	/* ignoring second returned value */
	q1, _ := divide(18, 3)
	fmt.Println("Quotient:", q1)

	checkAccess("Ritika", false)
}
