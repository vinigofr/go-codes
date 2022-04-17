package main

import (
	"fmt"
	"sync"
)

var theNumber int
var sw sync.WaitGroup
var mu sync.Mutex

func main() {
	sw.Add(2)
	go hahaha()
	go hahaha()
	sw.Wait()
}

func hahaha() {
	for i := 0; i < 5; i++ {
		// mu.Lock()
		theNumber += 1
		fmt.Println(theNumber)
		// mu.Unlock()
	}
	sw.Done()
}