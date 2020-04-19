package cc150

import "strings"

//题目
//给定一个字符串,编写一个函数判定其是否为某个回文串的排列之一。回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。回文串不一定是字典当中的单词。
//
//示例: 输入: Tact Coa
//
//输出: True (排列有 "taco cat" 、 "atco cta" ,等等)
//注意, 忽略空格和大小写

func ReverseString(s0 string, s1 string) bool {
	// First of all, remove all blank space and convert from uppercase/lowercase to lowercase
	s00 := make([]byte, 0)
	for i:=0; i<len(s0); i++ {
		if s0[i] == 32 {
			continue
		}
		s00 = append(s00, s0[i])
	}
	s000 := strings.ToUpper(string(s00))
	s11 := make([]byte, 0)
	for j:=0; j<len(s1); j++ {
		if s1[j] == 32 {
			continue
		}
		s11 = append(s11, s1[j])
	}
	s111 := strings.ToUpper(string(s11))
	if len(s000) != len(s111){
		return false
	}
	for i:=0; i<len(s000); i++ {
		j := len(s111)-1-i
		if s000[i] != s111[j] {
			return false
		}
	}
	return true
}