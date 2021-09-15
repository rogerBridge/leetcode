package leetcode

func longestCommonPrefix(strs []string) string {
	minLength := len(strs[0])
	for _, v := range strs {
		if len(v) < minLength {
			minLength = len(v)
		}
	}
	common := make([]byte, 0, minLength)
	for i := 0; i < minLength; i++ {
		compare := strs[0][i]
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != compare {
				return string(common)
			}
		}
		common = append(common, compare)
	}
	return string(common)
}
