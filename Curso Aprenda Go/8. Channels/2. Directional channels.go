// Quando utilizamos channels, podemos ter canais receivers e senders.
// Ou seja, enviar ou receber.

// Sintaxe:
// Para adicionar dados (send)
// channel <-[valor]

// Para remover dados (receive)
// <- channel
package main

import "fmt"

func main() {
	// Canal bidirecional.
	canal := make(chan int)

	go send(canal)
	receive(canal)
}

// A função send() recebe um canal como parâmetro.
// Ele armazena um int de valor 42 na variável do tipo canal chamada "s".
func send(s chan<- int) {
	fmt.Printf("%T\n", s)
	s <- 42 // Aqui atriuimos o valor 42
}

// A função receiver jogará o valor de "s" para "r"
// Assim, teremos em um novo canal o valor que antes, estava em em send()
func receive(r <-chan int) {
	fmt.Printf("%T\n", r) // Imprimindo seu tipo

	// Removemos o valor e o imprimimos
	fmt.Println("O valor recebido do canal foi", <- r)

	// Caso tentarmos adicionar valor a um canal que é apenas receiver
	// não fucionará
	// Ex: r <- 5

	// Mensagem de erro:
	// invalid operation: cannot send to receive-only channel r
	// (variable of type <-chan int)compilerInvalidSend
}