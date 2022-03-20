package main

// Imutais
// Podem ter tipos ou não
// As constantes não tipadas só terão um tipo quando realmente forem usadas.

// Declarando várias constantes de uma única vez
const (
	nome = "vinicius"
	sobrenome = "gouveis"
	idade = 22
)

// A mesma coisa acontece para variáveis
var (
	a = "vinicius"
	b = "gouveis"
	c = 22
)

const valorX = 10 // Este valor, apesar de ter um inteiro dentro de si, só terá um tipo quando for usado.

var valorY = 10 // Este valor já recebe um tipo na execução do código.
var valorZ float64
func main() {

	// Se eu tentar atribuir valorX para valorZ, dará certo pois valorX ainda não tem tipo.

	// Se eu tentar atribuir valorY para valorZ, dará errado pois valorY já tem tipo atribuido a si, que é float64.

	// valorZ = valorX - Dá certo
	// valorZ = valorY - Dá errado.
}