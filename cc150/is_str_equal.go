package cc150


// 给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串
// 其实还是看分布是否相等, map是否相等或者[]int是否相等
// 1. 假设两个被比较的字符串都是ascii, ascii可以显示的32~128, 我们就给128长度的[]int
func IsStrEqual(str0 string, str1 string) bool {
	if len(str0) == 0 && len(str1) == 0 {
		return true
	}
	if len(str0) == 0 || len(str1) == 0{
		return false
	}
	if len(str0) != len(str1) {
		return false
	}
	frequency1 := make([]uint8, 128)
	frequency2 := make([]uint8, 128)
	for i:=0; i<len(str0); i++ {
		frequency1[uint8(str0[i])] += 1
	}
	for j:=0; j<len(str1); j++ {
		frequency2[uint8(str1[j])] += 1
	}
	// compare two slice is equal
	for i:=0; i<len(str0); i++ {
		if frequency1[i] != frequency2[i] {
			return false
		}
	}
	return true
}

// If two string that compared is not only ASCII, we need map to judge if two str is equal
func IsStrEqualUTF8(str0 string, str1 string) bool {
	if len(str0)==0 && len(str1)==0 {
		return true
	}
	if len(str0) == 0 || len(str1) == 0 {
		return false
	}
	if len(str0) != len(str1) {
		return false
	}
	m0 := make(map[int32]int)
	m1 := make(map[int32]int)
	for i:=0; i<len(str0); i++ {
		if _, ok := m0[int32(str0[i])]; !ok {
			m0[int32(str0[i])] = 0
			m0[int32(str0[i])] += 1
		}
		m0[int32(str0[i])] += 1
	}
	for j:=0; j<len(str1); j++{
		if _, ok := m1[int32(str1[j])]; !ok {
			m1[int32(str0[j])] = 0
			m1[int32(str0[j])] += 1
		}
		m1[int32(str1[j])] += 1
	}
	// compare two map is equal
	for k, _:= range m0 {
		if m0[k] != m1[k] {
			return false
		}
	}
	return true
}