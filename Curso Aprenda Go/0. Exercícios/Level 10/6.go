package main

import "fmt"

var canalGlobal = make(chan int)

func main() {
	for i := 0; i < 10; i++ {
		go enviaAoCanal()
	}

	for i := 0; i < 50; i++ {
		fmt.Println(<-canalGlobal)
	}
}

func enviaAoCanal() {
	for i := 0; i < 5; i++ {
		canalGlobal <- i
	}
}