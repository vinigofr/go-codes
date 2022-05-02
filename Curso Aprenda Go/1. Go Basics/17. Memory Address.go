package main

import "fmt"

func main() {
	var name = "Vinicius"

	fmt.Println("Referenciando nome", &name, "<- Local em memória")

	// Pegamos o &name, que é o local de memória e alocamos ao 
	var name2 *string = &name 
	fmt.Println("Desreferenciando nome", *name2, "<- Valor")
}