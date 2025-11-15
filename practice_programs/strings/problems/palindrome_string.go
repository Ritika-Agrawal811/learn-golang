package problems

import "fmt"

/*
here reversed += string(str[i]) creates a new string everytime in golang

Go creates a new string, copies the old contents, appends the new character,
and then returns the new string. Since strings are immutable in Go,
every += is a new allocation, and that can get expensive in loops.

This is okay for numbers as Integers in Go are mutable primitives, so updating
reversed doesn’t allocate memory repeatedly — it's fast and cheap.

Time complexity is O(n²) because of repeated string concatenation
Space complexity is O(n) (because reversed string is length n)
*/
func IsPalindrome(str string) bool {

	reversed := ""

	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	return reversed == str
}

/* Time complexity is ---> O(n) n is the no. of chars in a string */
func OptimizedPalindromeCheck(str string) bool {

	/* loop till half the string */
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}

func RunPalindromeStringTest() {
	cases := []string{"aaa", "abc", "beam", "naman"}

	for _, str := range cases {
		result := OptimizedPalindromeCheck(str)
		fmt.Printf("%s is a palindrome? %v \n", str, result)
	}
}
