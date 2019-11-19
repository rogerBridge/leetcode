package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	sList := strings.Split(s, " ")
	r := make([]string, 0, 100)
	for _, v := range sList {
		if len(v) > 0 {
			r = append(r, v)
		}
	}
	//for _, v := range r {
	//	fmt.Print(v)
	//}
	return len(r[len(r)-1])
}

//func isLetter(s string) bool {
//	if unicode.IsLetter()
//}