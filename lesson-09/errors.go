package main

import (
	"errors"
	"fmt"
)

// Function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Error: cannot divide by zero
	} else {
		fmt.Println("Result:", result)
	}
}
