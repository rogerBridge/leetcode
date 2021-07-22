package sword

// 将一个字符串中的空格替换成 "%20"。
func ReplaceSpaces(s string) string {
	ss := ""
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		if s[i] == 32 {
			//log.Printf("Spaces Arise!")
			ss += "%20"
		} else {
			ss += string(s[i])
		}
	}
	return ss
}
