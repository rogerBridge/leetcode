package cc150_test

import (
	"fmt"
	"leetcode/cc150"
	"testing"
)

func TestIsStrEqual(t *testing.T) {
	testLists := [][]string{
		[]string{"", "", "true"},
		[]string{"hello", "eholl", "true"},
		[]string{"", "h", "false"},
		[]string{"1234", "4231", "true"},
	}
	for _, v := range testLists {
		if cc150.IsStrEqual(v[0], v[1]) == transStr2Bool(v[2]) {

		} else {
			t.Errorf("args: %s, %s, result: %s", v[0], v[1], v[2])
			fmt.Println("Fail")
		}
	}
}

func transStr2Bool(s string) bool {
	if s == "true" {
		return true
	}
	if s == "false" {
		return false
	}
	fmt.Println("转换过程中出现错误", s)
	return false
}
