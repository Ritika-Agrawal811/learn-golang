package problems

import (
	"fmt"
	"strings"
	"unicode"
)

/*
Time -> O(n * m)
Space -> O(n)

bug -- doesn't account for empty banned array
*/
func FindFrequentWord(paragraph string, banned []string) string {
	count := make(map[string]int)

	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	for _, word := range words {
		lowercase := strings.ToLower(word)
		for _, bannedWord := range banned {
			if lowercase == strings.ToLower(bannedWord) {
				break
			} else {
				count[lowercase]++
			}
		}
	}

	maxCount := 0
	maxWord := ""

	for word, freq := range count {
		if freq > maxCount {
			maxWord = word
			maxCount = freq
		}
	}

	return maxWord
}

/*
fixed and optmized version

Time -> O(n)
Space -> O(n)
*/
func OptimizedFindFrequentWord(paragraph string, banned []string) string {
	freq := make(map[string]int)

	bannedSet := make(map[string]bool)

	for _, word := range banned {
		bannedSet[strings.ToLower(word)] = true
	}

	fmt.Printf("banned set is %v\n", bannedSet)

	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	fmt.Printf("words is %v\n", words)

	for _, word := range words {
		lowercase := strings.ToLower(word)
		if !bannedSet[lowercase] {
			freq[lowercase]++
		}
	}

	fmt.Printf("freq is %v\n", freq)

	maxCount := 0
	maxWord := ""

	for word, count := range freq {
		if count > maxCount {
			maxWord = word
			maxCount = count
		}
	}

	return maxWord
}

func RunFindFrequentWordTest() {
	paragraph := "a, a, a, a, b,b,b,c, c"
	banned := []string{"a"}

	response := OptimizedFindFrequentWord(paragraph, banned)
	fmt.Printf("frequent not banned word is %s", response)
}
