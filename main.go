package main

import (
	"fmt"
	"leetcode/others"
	"leetcode/sword"
	"log"
	"runtime"
	"time"
)

func main() {
	fmt.Println(others.IsMatchString("ffabcdef", "abc"))
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