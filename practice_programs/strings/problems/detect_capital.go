package problems

import (
	"fmt"
	"unicode"
)

/*
Time cmplexity -> O(n)
Space -> O(1)
*/
func DetectCapital(word string) bool {
	capitals := 0

	for _, char := range word {
		if unicode.IsUpper(char) {
			capitals++
		}
	}

	if capitals == 1 && unicode.IsUpper(rune(word[0])) {
		return true
	}

	return capitals == len(word) || capitals == 0

}

/* could reduce some time like this */
func OptmizedDetectCapital(word string) bool {
	if len(word) == 1 {
		return true
	}

	firstUpper := unicode.IsUpper(rune(word[0]))
	secondUpper := unicode.IsUpper(rune(word[1]))

	// Case 1: If first is lowercase, all must be lowercase
	if !firstUpper && secondUpper {
		return false
	}

	// Case 2: From index 1 onward, all must match the 2nd letter's case
	for _, ch := range word[2:] {
		if unicode.IsUpper(ch) != secondUpper {
			return false
		}
	}

	return true
}

func RunDetectCapitalTest() {
	word := "eLeetcode"
	response := DetectCapital(word)

	fmt.Printf("is valid capital? %v", response)
}
