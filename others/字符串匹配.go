package others

// IsMatchString   观察subStr是否被包含进了str之中, 如果包含在其中, 返回subStr在Str中的下标, 否则返回-1
func IsMatchString(str string, subStr string) int {
	for i := 0; i < len(str)-len(subStr)+1; i++ {
		if str[i:i+len(subStr)] == subStr {
			return i
		}
	}
	return -1
}
