package main

import "fmt"

func main() {
	channelA := make(chan int)
	channelB := make(chan int)

	x := 10 

	go func(x int) {
		for i := 0; i < x; i++ {
			channelA <- i
		}
	}(x/2)

	go func(x int) {
		for i := 0; i < x; i++ {
			channelB <- i
		}
	}(x/2)

	// É como um switch case para canais. Quando um canal é acionado,
	// entra no case específico
	for i := 0; i < x; i++ {
		select {
		case v := <-channelA:
			fmt.Println("Recebeu do canal A:", v)
		case v := <-channelB:
			fmt.Println("Recebeu do canal B:", v)
		}
	}
}
 