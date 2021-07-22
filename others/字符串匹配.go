package others

// 观察subStr是否被包含进了str之中
func IsMatchString(str string, subStr string) int {
	for i:=0; i<len(str)-len(subStr)+1; i++ {
		if str[i:i+len(subStr)] == subStr {
			return i
		}
	}
	return -1
}
