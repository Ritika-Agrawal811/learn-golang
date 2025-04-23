package numbers

import (
	"fmt"
	"math"
)

func CountDigits(num int) int {
	if num == 0 {
		return 1
	}

	if num < 0 {
		num = -num
	}

	return int(math.Log10(float64(num))) + 1
}

func RunCountDigitsTest() {
	cases := []int{1222, 567, 900, 1, -4555}

	for _, num := range cases {
		count := CountDigits(num)
		fmt.Printf("digits count for %d is %d\n", num, count)
	}
}
