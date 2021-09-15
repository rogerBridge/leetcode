package main

import (
	"fmt"
	"leetcode/leetcode"
)

func main() {
	// traceMemStats()
	// timeTrack()
	//v := leetcode.LengthOfLongestSubstring("abcdea")
	//fmt.Println(sword.Bubble([]int{11, 3, 5, 2, 5, 8, 7}))
	//fmt.Println(256 & (256-1))
	//fmt.Println(sword.IfSquareOfTwo(65536))
	leetcode.IsValid("){")
}

// func traceMemStats() {
// 	var ms runtime.MemStats
// 	runtime.ReadMemStats(&ms)
// 	log.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
// }

// func timeTrack() {
// 	start := time.Now()
// 	log.Println(sword.ReplaceSpaces("hello world"), "elapsed:", time.Since(start))
// }

func printASCII() {
	s := "安瑞gg峰"
	for i, v := range s {
		fmt.Printf("%d: %x\n", i, rune(v))
	}
	fmt.Println([]byte(s))
	fmt.Println([]rune(s))
	fmt.Printf("%b\n", 229)
}
