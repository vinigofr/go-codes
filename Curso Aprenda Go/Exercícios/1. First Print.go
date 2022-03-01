// Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores x, y e z.

package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	stringResult := fmt.Sprintln(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(stringResult)
}