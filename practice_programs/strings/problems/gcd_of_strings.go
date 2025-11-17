package problems

import (
	"fmt"
)

/*
Time -> O(n + m)
Space -> O(n + m)

n -- length of str1
m -- length of str2
*/
func GcdOfStrings(str1 string, str2 string) string {

	if str1+str2 != str2+str1 {
		return ""
	}

	n := len(str1)
	m := len(str2)

	g := gcd(m, n)

	// length := n
	// if m < n {
	// 	length = m
	// }

	// for i := 2; i <= length; i++ {
	// 	if n%i == 0 && m%i == 0 {
	// 		gcd = i
	// 	}
	// }

	return str1[:g]
}

/* improved gcd calculation using Euclid’s algorithm */
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func RunGcdOfStringsTest() {
	// str1 := "ABCDEF"
	// str2 := "ABC"

	str1 := "ABABABABAB"
	str2 := "ABABAB"

	response := GcdOfStrings(str1, str2)
	fmt.Printf("gcd is %s", response)
}
