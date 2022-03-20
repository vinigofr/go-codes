package main

import "fmt"

// Para fazer o deslocamento de bits em go, é nesessário usar um dos operadores:
// esquerda << ou >> direita.



func main() {
	x := 2
	y := x << 1
	z := x >> 1

	fmt.Printf("%d Valor original, em binário, %b\n", x, x)
	fmt.Printf("%b << Bit deslocado para a esquerda\n", y)
	fmt.Printf("%b >> Bit deslocado para a direita", z)

}