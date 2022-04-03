package main

import "fmt"

// Falando de grosso modo, o iota é como se fosse um multiplicador de coisinhas
// Pod ser declarado igual a forma abaixo
const (
	a = iota
	b = iota
	c = iota
)

// E também assim:
const (
	d = iota * iota // Formula
	_ // Valores podem ser omitidos com o _ underscore
	f
	g
	h
)

func main() {
	fmt.Print(a, b, c, "\n")
	fmt.Print(d, f, g, h)

}