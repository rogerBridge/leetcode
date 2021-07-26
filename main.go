package main

import (
	"fmt"
	"leetcode/sword"
	"log"
	"runtime"
	"time"
)

func main() {
	// fmt.Println(others.IsMatchString("ffabcdef", "abc"))
	// log.Println(binarysearch.Sqrt(10))
	// priceList := []int{1, 2, 3, 11, 8, 9}
	// leetcodeofficial.MaxProfit(priceList)
	// traceMemStats()
	// timeTrack()
	printASCII()
}

func traceMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func timeTrack() {
	start := time.Now()
	log.Println(sword.ReplaceSpaces("hello world"), "elapsed:", time.Since(start))
}

func printASCII() {
	s := "安瑞gg峰"
	for i, v := range s {
		fmt.Printf("%d: %x\n", i, rune(v))
	}
	fmt.Println([]byte(s))
	fmt.Printf("%b\n", 229)
}
