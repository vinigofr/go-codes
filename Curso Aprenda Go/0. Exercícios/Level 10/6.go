package main

import "fmt"

var canalGlobal = make(chan int)

func main() {
	for i := 0; i < 10; i++ {
		go enviaAoCanal(chan<- int (canalGlobal))
	}

	for i := 0; i < 50; i++ {
		fmt.Println(<-canalGlobal)
	}
}

func enviaAoCanal(chan<- int) {
	for i := 0; i < 5; i++ {
		canalGlobal <- i
	}
}
