package cc150_test

import (
	"fmt"
	"leetcode/cc150"
	"testing"
)

func TestCompressString(t *testing.T) {
	testLists := [][]string{
		[]string{"aaabbcccddee", "a3b2c3d2e2"},
		[]string{"aaabbcccddaaa", "a3b2c3d2a3"},
	}
	for _, v := range testLists {
		if cc150.CompressString(v[0]) == v[1]  {

		} else {
			t.Errorf("args: %s, actual: %v, wanted: %v", v[0], cc150.CompressString(v[0]), v[1])
			fmt.Println("Fail")
		}
	}
}
