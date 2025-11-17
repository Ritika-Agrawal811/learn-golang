package problems

import "fmt"

/*
Time complexity is O(n) -- each index is visited at most once between the outer and inner loop combined.
Space -> O(1)

this is known as two-pointer technique
*/
func FindLargeGroups(s string) [][]int {
	groups := [][]int{}

	i := 0

	for i < len(s)-2 {
		if s[i] != s[i+1] || s[i] != s[i+2] {
			i++
			continue
		}

		j := i + 1

		for j < len(s) {
			if s[j] != s[i] {
				break
			}
			j++
		}

		groups = append(groups, []int{i, j - 1})
		i = j
	}

	return groups
}

/* improved version to reduce extra conditions */
func OptimizedFindLargeGroups(s string) [][]int {
	groups := [][]int{}

	i := 0

	for i < len(s) {
		j := i + 1

		for j < len(s) && s[j] == s[i] {
			j++
		}

		if j-i >= 3 {
			groups = append(groups, []int{i, j - 1})
		}

		i = j
	}

	return groups
}

func RunFindLargeGroupsTest() {
	s := "abbxxxxzzy"
	response := OptimizedFindLargeGroups(s)
	fmt.Println(response)
}
