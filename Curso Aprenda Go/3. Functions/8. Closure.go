package main

import "fmt"

// Quando utilizamos o chamad "closure", fazemos uma espécis de cerco.
// Conseguimos capturar um escopo específico para fazer seu uso válido.

func main() {
	a := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	// Um novo escopo é alocado na variável B
	b := i()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}