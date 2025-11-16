package problems

import (
	"fmt"
	"unicode"
)

/*
Time complexity -> O(n)
space -> O(n)
*/
func SwapLetters(s string) string {
	runes := []rune(s)

	left, right := 0, len(runes)-1

	for left < right {
		if !unicode.IsLetter(runes[left]) {
			left++
		} else if !unicode.IsLetter(runes[right]) {
			right--
		} else {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}

	return string(runes)
}

func RunSwapLettersTest() {
	s := "bn-jh"
	response := SwapLetters(s)
	fmt.Printf("swapped strin is : %s", response)
}
