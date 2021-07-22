package main

import (
	"log"
	"runtime"
	"time"
)

// global variables
var cpuNum = runtime.NumCPU()

func init() {
	log.Printf("Program init :)")
	runtime.GOMAXPROCS(cpuNum)
	log.Printf("CPU's num is: %d", cpuNum)
}

func main() {
	t0 := time.Now()
	log.Printf("%v", t0)
	for i:=0; i<cpuNum; i++ {
		go singleThreadFull()
	}
	time.Sleep(1000*time.Second)
	log.Printf("%v", time.Since(t0))
}

func singleThreadFull() {
	//for {
	//	for i := 0; i<4e7; i++ {
	//	}
	//	time.Sleep(10*time.Millisecond)
	//}
	for i:=0;;{
		i++
		//time.Sleep(250*time.Nanosecond)
	}
}
