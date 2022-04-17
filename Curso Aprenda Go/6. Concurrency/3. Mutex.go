package main

// O mutex faz que uma variável fique trancada
// Sendo apenas ser possível seu acesso através de uma
// única Thread.

// Dessa forma, conseguimos eliminar as race conditions
import (
	"fmt"
	"sync"
)

var theNumber int
var sw sync.WaitGroup
var mu sync.Mutex

func main() {
	sw.Add(5)
	go hahaha()
	go hahaha()
	go hahaha()
	go hahaha()
	go hahaha()
	sw.Wait()
}

func hahaha() {
	for i := 0; i < 5; i++ {
		mu.Lock()
		theNumber += 1
		fmt.Println(theNumber)
		mu.Unlock()
	}
	sw.Done()
}