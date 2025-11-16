package problems

import (
	"fmt"
	"unicode"
)

/*
Time complexity -> O(n)
space -> O(1)
*/
func CheckValidWord(word string) bool {
	if len(word) < 3 {
		return false
	}

	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	vowelsCount, consonantsCount := 0, 0

	for _, char := range word {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}

		if unicode.IsLetter(char) {
			if vowels[char] {
				vowelsCount++
			} else {
				consonantsCount++
			}
		}
	}

	return true && vowelsCount > 0 && consonantsCount > 0
}

func RunCheckValidWordTest() {
	word := "234Adas"
	response := CheckValidWord(word)

	fmt.Printf("is word valid? %v", response)
}
