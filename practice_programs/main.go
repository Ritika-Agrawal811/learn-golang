package main

import (
	"github.com/Ritika-Agrawal811/learn-golang/practice_programs/miscellaneous"
	"github.com/Ritika-Agrawal811/learn-golang/practice_programs/numbers"
	"github.com/Ritika-Agrawal811/learn-golang/practice_programs/strings"
)

func main() {

	/* Palindrome test */
	miscellaneous.RunPalinromeCheckTest()

	/* Roman to Integer */
	miscellaneous.RunRomanToNumberTest()

	/* Longest common prefix */
	miscellaneous.RunCommonPrefixTest()

	/* numbers programs */
	/* prime check */
	numbers.RunPrimeCheckTest()

	/* factorials */
	numbers.RunFactorialTest()

	/* sum of digits */
	numbers.RunSumOfDigitsTest()

	/* reverse a number */
	numbers.RunReversedNumberTest()

	/* armstrong number */
	numbers.RunArmstrongNumberTest()

	/* count the digits */
	numbers.RunCountDigitsTest()

	/* strings programs */
	/* check a palindrome string */
	strings.RunPalindromeStringTest()

}
