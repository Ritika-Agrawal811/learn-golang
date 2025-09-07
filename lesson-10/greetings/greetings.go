package greetings

import "fmt"

/*
 A simple function to include a person's name
 in a greeting
*/
func Hello(name string) string {
	message := fmt.Sprintf("Hello, %s. Welcome!", name)

	return message
}