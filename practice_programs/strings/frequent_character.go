package strings

import "fmt"

/*
Time Complexity is ---> O(n)

	Space Complexity is ---> O(k) where k is the no of unique letters
*/
func FindFrequentCharacter(str string) string {
	letters := make(map[rune]int)

	for _, char := range str {
		if count, found := letters[char]; found {
			letters[char] = count + 1
		} else {
			letters[char] = 1
		}
	}

	fmt.Println(letters)

	count := 0
	var freqChar rune
	for char, freq := range letters {
		if freq > count {
			count = freq
			freqChar = char
		}
	}

	return string(freqChar)
}

/*
we can use just one loop to do both. Complexity is still same
*/
func OptimizedFindFrequentCharacter(str string) string {
	letters := make(map[rune]int)

	count := 0
	var freqChar rune

	for _, char := range str {
		letters[char]++ // directly increment the count

		if letters[char] > count {
			count = letters[char]
			freqChar = char
		}
	}

	return string(freqChar)
}

func RunFindFrequentCharacterTest() {
	cases := []string{"hello", "banana"}

	for _, word := range cases {
		result := FindFrequentCharacter(word)
		fmt.Printf("word: %s, most occurring letter: %s\n", word, result)
	}
}
