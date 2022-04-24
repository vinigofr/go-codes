package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup;

// É uma forma de de programar
func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)


	go send(par, impar)
	go receive(par, impar, converge)

	for v := range converge {
		fmt.Println("O valor recebido é igual a", v)
	}
}

func send(par, impar chan int) {
	x := 10
	for n := 0; n < x; n++ {
		if n % 2 == 0 {
			par <- n
		} else {
			impar <- n
		}
	}
	close(par)
	close(impar)
}

func receive(par, impar, converge chan int) {
	waitGroup.Add(2)
	go func(){
		for v := range par {
			converge <- v
		}
		waitGroup.Done()
	}()

	go func(){
		for v := range impar {
			converge <- v
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
	close(converge)
}