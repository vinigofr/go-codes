package main

import "fmt"

// Vamos criar um tipo chamado "hotdog"

type hotdog int
var b hotdog

func main() {
	b = 50
	// Imprimindo o tipo da variável, que é hotdog
	fmt.Printf("%T\n", b)

	// Modificando o tipo de uma variável já crianda
	x := int(b)

	fmt.Println(x)
}
