package main

import (
	"fmt"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//
//所有输入只包含小写字母 a-z 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-common-prefix
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func longestCommonPrefix(strs []string) string {
	// 首先, 找出字符串切片中, 长度最小的字符串的长度
	minStrLen := len(strs[0])
	for _, v := range strs {
		if minStrLen>len(v){
			minStrLen = len(v)
		}
	}
	fmt.Println("字符串数组中长度最小的字符串是:", minStrLen)
	// 然后, 依据最小字符串长度来查找公共前缀
	commonTime := 0
	I:
	for i:=0; i<minStrLen; i++{
		for j:=0; j<len(strs)-1; j++ {
			if strs[j][i] != strs[j+1][i]{
				break I
			}
		}
		commonTime += 1
	}
	return strs[0][:commonTime]
}

func main() {
	r := longestCommonPrefix([]string{"github", "gitlab", "gitt", "gittttt"})
	fmt.Println(r)
}