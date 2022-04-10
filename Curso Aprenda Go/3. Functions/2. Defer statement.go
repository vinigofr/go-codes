package main

import "fmt"

// O defer é um argumento que diz ao compilador
// que um específico trecho de código deve ser executado por último

// A ordem de defer importa.
// O primeiro a receber o sempre será o último

// É como se fosse uma pilha (FIFO) LILO

func main() {
	defer fmt.Println("Veio por último")
	fmt.Println("Veio primeiro")
	defer fmt.Println("Está no meio")
}