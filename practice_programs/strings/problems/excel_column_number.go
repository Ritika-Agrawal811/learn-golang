package problems

import (
	"fmt"
)

/*
we are building in opposite to avoid math.Pow()

like in base 10 we have --

	567
	= 5×10² + 6×10¹ + 7×10⁰

	Build iteratively:
	result = 0
	result = 0*10 + 5 = 5
	result = 5*10 + 6 = 56
	result = 56*10 + 7 = 567

Similarly for base 26 --

	"ABC"
	= A×26² + B×26¹ + C×26⁰

	Build iteratively:
	result = 0
	result = 0*26 + A = 1
	result = 1*26 + B = 28
	result = 28*26 + C = 731
*/
func FindExcelColumn(column string) int {
	sum := 0

	for i := 0; i < len(column); i++ {
		digit := int(column[i]-'A') + 1
		sum = sum*26 + digit
	}

	return sum
}

func RunFindExcelColumnTest() {
	column := "AB"
	response := FindExcelColumn(column)
	fmt.Printf("excel column number is %d", response)
}
