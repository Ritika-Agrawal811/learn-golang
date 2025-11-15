package problems

import "fmt"

/*
Time complexity -> O(n * m)
Space -> O(1)
*/
func FindFirstOccurence(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	// try each possible starting index
	limit := len(haystack) - len(needle)

	// first loop is for possible indices
	for i := 0; i <= limit; i++ {
		matched := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				matched = false
				break
			}
		}
		if matched {
			return i
		}
	}
	return -1
}

func RunFindFirstOccurenceTest() {
	needle := "issip"
	haystack := "mississippi"
	response := FindFirstOccurence(haystack, needle)

	fmt.Printf("first occurence of %s in %s is %d", needle, haystack, response)
}
