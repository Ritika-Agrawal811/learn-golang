package strings

import (
	"fmt"
	"strings"
	"unicode"
)

/*
Time complexity is ---> O(n)
Space complexity here is also O(n) because we are converting string to lowercase
*/
func CountVowelsAndConsonants(str string) {
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}

	vowelsCount := 0
	consonantsCount := 0

	for _, char := range strings.ToLower(str) {

		/* skip the non-letters */
		if !unicode.IsLetter(char) {
			continue
		}

		if vowels[char] {
			vowelsCount++
		} else {
			consonantsCount++
		}
	}

	fmt.Printf("word: %s, vowels: %d, consonants: %d\n", str, vowelsCount, consonantsCount)
}

/*
reduce space from creating a lower case copy by converting each rune to lowercase in the loop

	Space complexity reduces to O(1)
*/
func OptimizedCount(str string) {
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}

	vowelsCount := 0
	consonantsCount := 0

	for _, char := range str {

		/* skip the non-letters */
		if !unicode.IsLetter(char) {
			continue
		}

		r := unicode.ToLower(char)

		if vowels[r] {
			vowelsCount++
		} else {
			consonantsCount++
		}
	}

	fmt.Printf("word: %s, vowels: %d, consonants: %d\n", str, vowelsCount, consonantsCount)
}

func RunCountVowelsAndConsonantsTest() {
	cases := []string{"heLLo"}

	for _, word := range cases {
		OptimizedCount(word)
	}
}
