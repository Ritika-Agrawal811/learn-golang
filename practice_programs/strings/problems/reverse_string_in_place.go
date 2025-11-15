package problems

import "fmt"

/*
time complexity -> O(n)
space -> O(1)
*/
func ReverseStringInPlace(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func RunReverseStringInPlaceTest() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Printf("reversed string is: %v", s)

	ReverseStringInPlace(s)

	fmt.Printf("reversed string is: %v", s)
}
