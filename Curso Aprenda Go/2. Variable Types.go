package main

import "fmt"

// Fora do escopo, deve-se fazer um dos seguintes casos:
const z int = 10
var u int = 10

// O gopher " := " é o símbolo de criação de variáveis em GO
// Em português é chamado de marmota.
// A tipagem é atribuída automaticamente.
func main() {
	// Atribuições
	x := 10
	y := "Olá, bom dia!"
	b := true
	f := 1.5

	fmt.Printf("x: %v, %T\n", x, x);
	fmt.Printf("y: %v, %T\n", y, y);
	fmt.Printf("b: %v, %T\n", b, b);
	fmt.Printf("f: %v, %T\n", f, f);

	// Reatribuições com um igual. " = "
	x = 20
	y = "Olá, boa tarde!"
	b = false
	f = 2.5

	// Sobre strings em Go
	fmt.Println(stringLiterals())
}

func stringLiterals()(string, string) {
	// Posso retornar tanto strings formatadas
	anyString := "Isso é uma string\n"

	// Ou strings literais
	anotherString := `Isso é outra string`

	// O Sprint() exibe a concatenação de strings
	fmt.Println(anyString, anotherString)

	return anyString, anotherString
}
