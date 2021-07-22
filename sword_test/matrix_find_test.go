package sword_test

import (
	"fmt"
	"leetcode/sword"
	"testing"
)

func TestEditOnce(t *testing.T) {
	testLists := [][][]int{
		[][]int{[]int{1,2,3}, []int{2,3,4}, []int{3,4,5}, []int{5}},
	}
	for _, v := range testLists {
		if sword.FindInMatrix(v[len(v)-1][0], v) == true  {

		} else {
			t.Errorf("args: %v,%v, actual: %v", v[0:len(v)-2], v[len(v)-1], sword.FindInMatrix(v[len(v)-1][0], v), )
			fmt.Println("Fail")
		}
	}
}
