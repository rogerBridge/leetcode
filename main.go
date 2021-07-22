package main

import (
	leetcodeofficial "leetcode/leetcodeOfficial"
	"leetcode/sword"
	"log"
	"runtime"
	"time"
)

func main() {
	// fmt.Println(others.IsMatchString("ffabcdef", "abc"))
	// log.Println(binarysearch.Sqrt(10))
	priceList := []int{1, 2, 3, 11, 8, 9}
	leetcodeofficial.MaxProfit(priceList)
	// traceMemStats()
	// timeTrack()
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
