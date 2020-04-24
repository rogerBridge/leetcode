package double_point

import "bytes"

//使用双指针，一个指针从头向尾遍历，一个指针从尾到头遍历，当两个指针都遍历到元音字符时，交换这两个元音字符。
//
//为了快速判断一个字符是不是元音字符，我们将全部元音字符添加到集合 HashSet 中，从而以 O(1) 的时间复杂度进行该操作。

func ReverseVowels(s string) string {
	ss := []uint8(s)
	i := 0
	j := len(ss) - 1
	vowels := []uint8{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for i < j {
		if bytes.Contains(vowels, []byte{ss[i]}) && bytes.Contains(vowels, []byte{ss[j]}) {
			// 字符串是不能变更的
			x := ss[i]
			ss[i] = ss[j]
			ss[j] = x
			i++
			j--
		} else if bytes.Contains(vowels, []byte{ss[i]}) == false && bytes.Contains(vowels, []byte{ss[j]}) {
			i++
		} else if bytes.Contains(vowels, []byte{ss[i]}) && bytes.Contains(vowels, []byte{ss[j]}) == false {
			j--
		} else {
			i++
			j--
		}
	}
	return string(ss)
}
