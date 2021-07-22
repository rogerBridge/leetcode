package cc150

import "math"

//字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。给定两个字符串,编写一个函数判定它们是否只需要一次(或者零次)编辑。
//
//示例:
//
//pale, ple -> true
//
//pales, pale -> true
//
//pale, bale -> true
//
//pale, bake -> false

// 插入, 删除, 替换
// 假设字符全部都是ASCII
func EditOnce(str0 string, str1 string) bool {
	if str0 == "" && str1 == "" {
		return true
	}
	s0 := make([]byte, 128)
	s1 := make([]byte, 128)
	for i:=0; i<len(str0); i++ {
		s0[str0[i]] += 1
	}
	for i:=0; i<len(str1); i++ {
		s1[str1[i]] += 1
	}
	// compare two slice
	for i:=0; i<len(s0); i++ {
		if math.Abs(float64(s0[i]) - float64(s1[i])) == 1.00 {
			return true
		}
	}
	return false
}