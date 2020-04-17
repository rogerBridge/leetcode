package cc150_test

import (
	"fmt"
	"leetcode/cc150"
	"testing"
)

func TestURLformat(t *testing.T) {
	testLists := [][]string{
		[]string{"asdf a", "asdf%20a"},
		[]string{"bsd a", "bsd%20a"},
		[]string{"abcdef g", "abcdef g"}, // 故意搞错测试用例
	}
	for _, v := range testLists {
		if cc150.URLformat(v[0]) == v[1] {

		} else {
			t.Errorf("args: %s, expect: %s, actual: %v", v[0], v[1], cc150.URLformat(v[0]))
			fmt.Println("Fail")
		}
	}
}
