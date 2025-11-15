package problems

import "fmt"

/*
Time complexity is O(n²) and space complexity is O(n)
because of repeated string concatenation
*/
func ReverseString(str string) string {

	reversed := ""

	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	return reversed
}

/* Time complexity --> O(n) and space complexity ---> O(n) */
func OptimizedReverseString(str string) string {
	runes := []rune(str) // converting the string to a slice of runes
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}

	return string(runes) // converting a slice of rune to a string.
}

func RunReverseStringTest() {
	cases := []string{"hello", "blah", "world"}

	for _, word := range cases {
		result := ReverseString(word)
		fmt.Printf("word: %s, reversed: %s \n", word, result)
	}
}
