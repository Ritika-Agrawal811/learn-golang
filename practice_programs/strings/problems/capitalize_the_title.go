package problems

import (
	"fmt"
	"strings"
)

/*
time -> O(n)
space -> O(n)
*/
func CapitalizeTheTitle(title string) string {
	words := strings.FieldsFunc(title, func(r rune) bool {
		return r == ' '
	})

	for i, word := range words {
		if len(word) > 2 {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		} else {
			words[i] = strings.ToLower(word)
		}
	}

	return strings.Join(words, " ")
}

/*
here we use strings.Builder
time and space complexity stay the same,
it just avoids repeated copying when concatenating strings
*/
func OptimizedCapitalizeTheTitle(title string) string {
	var b strings.Builder
	words := strings.Fields(title)

	for i, w := range words {
		if len(w) <= 2 {
			w = strings.ToLower(w)
		} else {
			w = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}

		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(w)
	}

	return b.String()
}

func RunCapitalizeTheTitleTest() {
	title := "First leTTeR of EACH Word"
	response := CapitalizeTheTitle(title)

	fmt.Printf("title is %s", response)
}
