package cc150_test

import (
	"fmt"
	"leetcode/cc150"
	"testing"
)

func TestReverseString(t *testing.T) {
	testLists := [][]string{
		[]string{"asdf a", "afDsa"},
		[]string{"abcdEf  g", "gFedcba"},
	}
	for _, v := range testLists {
		if cc150.ReverseString(v[0], v[1]) == true {

		} else {
			t.Errorf("args: %s,%s, actual: %v", v[0], v[1], cc150.ReverseString(v[0], v[1]))
			fmt.Println("Fail")
		}
	}
}
