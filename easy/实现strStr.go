package main

func strStr(haystack string, needle string) int {
	// 若needle字符串存在于haystack中, 返回needle字符串出现的第一个位置, 若needle字符串不存在与haystack中, 则返回-1
	if len(needle) > len(haystack) {
		return -1
	} else if len(haystack) == len(needle) && haystack==needle {
		return 0
	} else {
		for i:=0; i<=len(haystack)-len(needle); i++ {
			if haystack[i:len(needle)+i] == needle {
				return i
			}
		}
		return -1
	}
}