package problems

import (
	"fmt"
)

/*
Time complexity -> O(n)
space -> O(n)
*/
func SwapVowels(str string) string {
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	runes := []rune(str)

	left, right := 0, len(runes)-1

	for left < right {
		if !vowels[runes[left]] {
			left++
		} else if !vowels[runes[right]] {
			right--
		} else {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}

	return string(runes)
}

func RunSwapVowelsTest() {
	str := "IceCreAm"
	response := SwapVowels(str)
	fmt.Printf("swapped vowels string is : %s", response)
}
