package cc150

import "strings"

//字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。

func rotateString(s1 string, s2 string) bool {
	// 一个旋转体的两倍必定包含另一个旋转体
	if len(s1) != len(s2) {
		return false
	}
	return strings.Contains(s2+s2, s1)
}