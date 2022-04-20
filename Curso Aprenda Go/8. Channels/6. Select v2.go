package main

import "fmt"

func main() {
	canal := make(chan int)
	quit := make(chan int)

	go recebeQuit(canal, quit)
	enviaParaCanal(canal, quit)

}
// A função enviaParaCanal manda o valor para cima na função recebeQuit
func recebeQuit(canal chan int, quit chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Recebido:", <- canal)
	}
	quit <- 0
}

func enviaParaCanal(canal chan<- int, quit chan int) {
	numeroQualquer := 1
	for {
		select {
		case canal <- numeroQualquer:
			numeroQualquer++
		case <- quit:
			return
		}
	}
}