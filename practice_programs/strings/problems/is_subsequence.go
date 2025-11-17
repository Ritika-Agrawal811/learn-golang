package problems

import "fmt"

/*
Time: O(|s| + |t|)
Space: O(1)
*/
func IsSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}

	i := 0
	j := 0

	for i < len(s) {

		for j < len(t) {
			if s[i] == t[j] {
				break
			}
			j++
		}

		if j == len(t) {
			return false
		}

		i++
		j++
	}

	return true
}

/*
Time: O(|s| + |t|)
Space: O(1)

here we only increase i when we find the match in t
Go allows to add conditions for both pointers in one loop
*/
func OptimzedIsSubsequence(s, t string) bool {
	i := 0 // pointer in s

	for j := 0; j < len(t) && i < len(s); j++ {
		if s[i] == t[j] {
			i++
		}
	}

	return i == len(s)
}

func RunIsSubsequenceTest() {
	s := "axc" // "aec"
	t := "ahbgdc"

	response := IsSubsequence(s, t)
	fmt.Printf("is subsequence? %v", response)
}
