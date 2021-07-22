package cc150

import (
	"strconv"
)

// 利用字符重复出现的次数,编写一种方法,实现基本的字符串压缩功能.
// 比如,字符串 aabcccccaaa 会变为 a2b1c5a3 。若“压缩”后的字符串没有变短,则返回原先的字符串。你可以假设字符串中只包含大小写英文字母(a 至 z)

func CompressString(s string) string {
	m := make([]int32, 128)
	r := ""
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 0 {
			m[s[i]] += 1
		}
		if i == len(s) -1 {
			r += string(s[i])
			r += strconv.Itoa(int(m[s[i]]))
		} else {
			if s[i+1] == s[i] {
				m[s[i]] += 1
			} else {
				r += string(s[i])
				r += strconv.Itoa(int(m[s[i]]))
				m[s[i]] = 0
			}
		}
	}
	// 判断自己所得字符串长度和源字符串长度对比, 若优于原字符串, 则返回自己压缩后的字符串
	if len(r) < len(s) {
		return r
	} else {
		return s
	}
	//for i, v := range m {
	//	if v != 0 {
	//		r += string(i)
	//		r += strconv.Itoa(int(v))
	//	}
	//}
}