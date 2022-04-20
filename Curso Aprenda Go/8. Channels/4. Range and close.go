package main

import "fmt"

func main() {
	channel := make(chan int)
	total := 10

	// Mesmo que minha função receba como parâmetro um canal do tipo sender
	// ainda podemos enviar um canal do tipo bidirecional para ele.
	go myLoop(channel, total)

	for v  := range channel {
		fmt.Println(v)
	}
}

// Eu tenho que colocar coisa no canal
func myLoop(send chan<- int, total int) {
	for i := 0; i < total; i += 1 {
		fmt.Println(i)
		send <- i
	}

	// Precisamos fechar o canal para que o range acima não fique "esperando"
	// pelo final do for loop.
	close(send)
}