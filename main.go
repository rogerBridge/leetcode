package main

import (
	"fmt"
	"leetcode/cc150"
)

func main() {
	fmt.Println(cc150.HasDiffChar("我爱你"))
	fmt.Println(cc150.HasDiffChar("我爱你, 我的家"))
	fmt.Println(cc150.HasDiffChar1("abcdefg"))
	fmt.Println(cc150.HasDiffChar1("abcdefga"))
}
