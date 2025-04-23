package miscellaneous

func IsParenthesisValid(s string) bool {

	stack := []rune{}

	matching := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 || matching[char] != stack[len(stack)-1] {
				return false
			}

			// pop from stack
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0
}
