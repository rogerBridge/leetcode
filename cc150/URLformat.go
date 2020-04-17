package cc150

//编写一个方法，将字符串中的空格全部替换为“%20”。假定该字符串尾部有足够的空间存放新增字符串，并且知道字符串的“真实”的长度。（注：用Java实现的话，请使用字符数组实现，以遍直接在数组上操作。）
//示例：
//输入：“Mr John Smith”, 13
//输出：“Mr%20John%20Smith”

func URLformat(s string) string {
	if len(s) == 0 {
		return ""
	}
	// convert before
	m := []byte(s)
	// convert after
	n := make([]byte, 0)
	for i:=0; i<len(m); i++ {
		if m[i] == 32 {
			n = append(n, '%', '2', '0')
		} else {
			n = append(n, m[i])
		}
	}
	return string(n)
}