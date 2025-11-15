package problems

import "fmt"

/*
Time complexity -> O(n)
space -> O(1)
*/
func LengthOfLastWord(word string) int {
	length := 0
	for i := len(word) - 1; i >= 0; i-- {
		if word[i] != ' ' {
			length++
		}

		if length != 0 && word[i] == ' ' {
			return length
		}
	}

	return length
}

func RunLengthOfLastWordTest() {
	word := "luffy is still joyboy"
	response := LengthOfLastWord(word)
	fmt.Printf("length of last word is %d", response)
}
