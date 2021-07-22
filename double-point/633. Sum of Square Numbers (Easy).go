package main

import "math"

//给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c。
//
//示例1:
//
//输入: 5
//输出: True
//解释: 1 * 1 + 2 * 2 = 5
// 
//
//示例2:
//
//输入: 3
//输出: False

// 还是左右两个指针, 左指针为0, 右指针为int(sqrt(c))
func JudgeSquareSum(c int) bool {
	i:=0
	j:= int(math.Sqrt(float64(c)))
	for ; i<=j; {
		if c == i*i + j*j {
			return true
		} else if c < i*i+j*j {
			j--
		} else{
			i++
		}
	}
	return false
}