package cc150

// 实现一个算法，确定一个字符串的所有字符是否全都不同。假使不允许使用额外的数据结构，又该如何处理？

// 假设这段字符串中有中文, 使用map, 统计频率
func HasDiffChar(tempStr string) bool {
	m := make(map[int32]int)
	for _, v := range tempStr {
		//if _, ok := m[v]; !ok {
		//	m[v] += 1
		//} else {
		//	m[v] += 1
		//}
		m[v] += 1
	}
	for _, v := range m {
		if v > 1 {
			return true
		}
	}
	return false
}

// 假设这个字符串时ascii, 那就直接上slice吧
func HasDiffChar1(tempStr string) bool {
	m := make([]int, 128)
	for _, v := range tempStr {
		m[int(byte(v))] += 1
	}
	for _, v := range m {
		if v > 1 {
			return true
		}
	}
	return false
}
