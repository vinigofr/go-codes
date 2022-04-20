package main

import "fmt"

// Os canais funcionam como aquelas corridas olimpicas que você entgrega o
// bastão a outros corredores, para que eles possam seguir seus caminhos.
// Ou seja, não adianta apenas criar um canal e atribuir um valor a ele
// se não podemos passá-lo adiante.
func main() {
	fmt.Println("Introdução")
	channel := make(chan int, 1) // Aqui determinamos o buffer de 1. 
	// Este canal suporta apenas um valor.
	// Para elevar isto, precisaríamos elevar seu valor para 2.
	
	// Uma go routine
	go func() {
		channel <- 42 // Atrubuindo um valor a um canal
	}()

	teste := <- channel // Podemos atribuir a uma variável
	fmt.Println(teste)


	// Em geral, não se usa buffer, mas ele existe...
	fmt.Println("Utilizando buffers")
	fmt.Println()
	channel2 := make(chan int, 5)
	channel2 <- 1
	channel2 <- 2
	channel2 <- 3
	channel2 <- 4
	channel2 <- 5

	fmt.Println(channel2)		 // Endereço de memória, ex: 0xc00007e000
	fmt.Println(<- channel2) // 1
	fmt.Println(<- channel2) // 2
	fmt.Println(<- channel2) // 3
	fmt.Println(<- channel2) // 4
	fmt.Println(<- channel2) // 5
}