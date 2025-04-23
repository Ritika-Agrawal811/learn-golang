package miscellaneous

import "fmt"

func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	commonPrefix := strs[0]

	for i := 1; i < len(strs); i++ {
		j := 0

		for j < len(commonPrefix) && j < len(strs[i]) && commonPrefix[j] == strs[i][j] {
			j++
		}

		commonPrefix = commonPrefix[:j]

		if commonPrefix == "" {
			return ""
		}
	}

	return commonPrefix
}

func RunCommonPrefixTest() {
	strs := []string{"flower", "flow", "flight"}

	prefix := LongestCommonPrefix(strs)
	fmt.Println(prefix)
}
