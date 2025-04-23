package miscellaneous

import "fmt"

func RomanToInteger(s string) int {
	romans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	prevChar := ""
	num := 0

	for i := len(s) - 1; i >= 0; i-- {
		currentChar := string(s[i])

		if (prevChar == "V" || prevChar == "X") && currentChar == "I" {
			num -= 1
		} else if (prevChar == "L" || prevChar == "C") && currentChar == "X" {
			num -= 10
		} else if (prevChar == "M" || prevChar == "D") && currentChar == "C" {
			num -= 100
		} else {
			num += romans[currentChar]
		}

		prevChar = currentChar

	}

	return num
}

// faster version using bytes
func ConvertRomanToInteger(s string) int {

	// bytes are written with single quotes
	romans := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	prevChar := 0
	num := 0

	for i := len(s) - 1; i >= 0; i-- {
		// for loop retrived bytes from strings hence no need to convert now
		currChar := romans[s[i]]

		if currChar < prevChar {
			num -= currChar
		} else {
			num += currChar
		}

		prevChar = currChar
	}

	return num

}

func RunRomanToNumberTest() {
	romanTestCases := []string{"MCMXCIV"}

	for _, num := range romanTestCases {
		result := ConvertRomanToInteger(num)
		fmt.Println(result)

	}
}
