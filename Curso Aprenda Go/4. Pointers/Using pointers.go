package main

import "fmt"

func main() {
	// Aqui temos uma variável
	x := 10
	fmt.Println(x)

	// Podemos armazenar o seu endereço da seguinte forma
	y := &x
	fmt.Println(y)

	// Podemos "desreferenciar" o endereço fazendo o seguinte
	z := *y
	fmt.Println(z)

	// Outra sintaxe maluca
	fmt.Println(*&x)

	// E os tipos
	fmt.Printf("%T %T\n", x, y) // tipo int &&  tipo ponteiro *int

	// Incrementando em ponteiros
	*y++
	fmt.Println(*y)

}