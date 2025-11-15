package problems

import "fmt"

/*
Time complexity -> O(n)
Space complexity -> O(n)
*/
func FindTheDifferentLetter(s string, t string) byte {
	letters := make(map[rune]int)

	for _, char := range s {
		letters[char]++
	}

	for _, char := range t {
		letters[char]--

		if letters[char] < 0 {
			return byte(char)
		}
	}

	return 0
}

/*
XOR cancels out all matching characters
Time complexity -> O(n)
Space complexity -> O(1)
*/

func OptimizedFindTheDifferentLetter(s string, t string) byte {
	var result byte = 0

	for _, char := range s {
		result ^= byte(char)
	}

	for _, char := range t {
		result ^= byte(char)
	}

	return result
}

func RunFindTheDifferentLetterTest() {
	s := "abcd"
	t := "abcde"

	response := OptimizedFindTheDifferentLetter(s, t)
	fmt.Printf("the different letter is %v", string(response))
}
