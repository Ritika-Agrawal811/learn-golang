package problems

import "fmt"

/*
Time complexity -> O(n * m)
Space is O(n)
*/
func CheckSubstringLetters(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	count := 0
	runes := []rune(magazine)

	for _, letter := range ransomNote {
		for i, char := range runes {
			if letter^char == 0 {
				count++
				runes[i] = ' '
				break
			}
		}
	}

	return count == len(ransomNote)
}

/*
Time complexity -> O(n + m)
Space is O(1) [ only 26 letters in alphabet ]
*/
func OptimizedCheckSubstringLetters(ransomNote string, magazine string) bool {
	freq := make(map[rune]int)

	if len(magazine) < len(ransomNote) {
		return false
	}

	for _, char := range magazine {
		freq[char]++
	}

	for _, char := range ransomNote {
		freq[char]--

		if freq[char] < 0 {
			return false
		}
	}

	return true
}

/*
Time complexity -> O(n + m)
Space is O(1) [ only 26 letters in alphabet ]

  - only byte operations
    # No heap allocations
    # No hashing
    # No map overhead
    # No Unicode decoding
    # Pure array indexing, which is CPU-cache friendly
    # Pure byte math
*/
func TrulyOptimizedVersion(ransomNote string, magazine string) bool {
	freq := [26]int{}

	if len(magazine) < len(ransomNote) {
		return false
	}

	for i := 0; i < len(magazine); i++ {
		freq[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		freq[ransomNote[i]-'a']--

		if freq[ransomNote[i]-'a'] < 0 {
			return false
		}
	}

	return true
}

func RunCheckSubstringLettersTest() {
	magazine := "kjj"
	ransomNote := "kk"

	response := TrulyOptimizedVersion(ransomNote, magazine)
	fmt.Printf("is possible? %v", response)
}
