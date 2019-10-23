package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)

const StrBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func init() {
	runtime.GOMAXPROCS(4)
}

func main() {
	concurrentNum := 1000
	result := make(chan int, concurrentNum)
	for i := 0; i < concurrentNum; i++ {
		go worker1(result)
	}
	go customer(result, concurrentNum)

	forever := make(chan bool)
	<-forever
	//_, _ = fmt.Scanln()
}

func worker(target chan int) int {
	log.Println(time.Now().UnixNano())
	var sum int
	for i := 0; i < 100000; i++ {
		rand.Seed(time.Now().UnixNano())
		sum += rand.Intn(2)
		//sum += rand.Int63() & 1
	}
	target <- sum
	return sum
}

func worker1(target chan int) {
	time.Sleep(2*time.Second)
	r := rand.Intn(2)
	target <- r
}

func customer(target chan int, concurrentNum int) {
	s := make([]int, concurrentNum)
	for i := 0; i < concurrentNum; i++ {
		s[i] = <-target
	}
	//for v := range target {
	//	s = append(s, v)
	//}
	fmt.Println(s)
	//for v := range target {
	//	fmt.Println(v)
	//}
}
