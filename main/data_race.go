package main

import (
	"fmt"
	"sync"
	"time"
)

var inited = false
var lock sync.Mutex

func Setup() {
	time.Sleep(time.Second)
	lock.Lock()
	inited = true
	lock.Unlock()
}

func Setup2() <-chan bool {
	time.Sleep(time.Second * 3)
	c := make(chan bool, 1)
	c <- true
	return c
}
func main() {
	//go Setup()
	//
	//for {
	//	lock.Lock()
	//	b := inited
	//	lock.Unlock()
	//	if b {
	//		break
	//	}
	//	time.Sleep(100 * time.Millisecond)
	//}

	//fmt.Println("setup succeed")

	if <-Setup2() {
		fmt.Println("setup succeed")
	}

}
