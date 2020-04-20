package cc150_test

import (
	"fmt"
	"leetcode/cc150"
	"testing"
)

func TestEditOnce(t *testing.T) {
	testLists := [][]string{
		[]string{"", "", "true"},
		[]string{"pale", "ple", "true"},
		[]string{"pales", "pale", "true"},
		[]string{"pale", "bale", "true"},
		[]string{"pale", "paless", "false"},
	}
	for _, v := range testLists {
		if cc150.EditOnce(v[0], v[1]) == TransStr2Bool(v[2])  {

		} else {
			t.Errorf("args: %s,%s, actual: %v, wanted: %v", v[0], v[1], cc150.EditOnce(v[0], v[1]), v[2])
			fmt.Println("Fail")
		}
	}
}