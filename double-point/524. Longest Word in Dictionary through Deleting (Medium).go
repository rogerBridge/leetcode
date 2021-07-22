package main

// 给定一个字符串和一个字符串字典，找到字典里面最长的字符串，该字符串可以通过删除给定字符串的某些字符来得到。如果答案不止一个，返回长度最长且字典顺序最小的字符串。如果答案不存在，则返回空字符串。

// 按照顺序, 把字典里面符合要求的子串添加到一个列表内, 找到字典序列最小, 长度最长的子串
func FindLongestWord(s string, d []string) string {
	r := make([]string, 0)
	for _, v := range d {
		if isSequence(s, v) {
			r = append(r, v)
		}
	}
	// 找出v之中最长的
	x := r[0]
	for _, v := range r {
		if len(v) > len(x) {
			x = v
		}
	}
	return x
}

// x字符是否包含子字符串y
func isSequence(x string, y string) bool {
	j:=0
	for i:=0; i<len(x) && j<len(y) ;i++  {
		if x[i] == y[j] {
			j++
		}
	}
	return j==len(y)
}