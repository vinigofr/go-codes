package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup
var contador int

func main() {
	waitGroup.Add(2)

	go thePrint()
	go thePrint()
	waitGroup.Wait()
}

func thePrint() {
	contador++
	fmt.Println("Go routine numero", contador)
	waitGroup.Done()
}