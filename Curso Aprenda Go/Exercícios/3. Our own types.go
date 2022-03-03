package main

import "fmt"

type myOwnType int

var x myOwnType
var y int

func main() {
	fmt.Printf(`
		Valor de x: %v
		Tipo de x: %T
	`, x, x)

	x = 42
	y = int(x)

	fmt.Printf(`
O novo valor de X é: %v
`, x)
	fmt.Printf(`O novo valor de Y é: %v`, y)
}