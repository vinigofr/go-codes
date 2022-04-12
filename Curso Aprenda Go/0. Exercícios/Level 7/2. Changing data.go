package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func mudeMe(p *pessoa) {
	// 4. O parâmetro vem na forma de ref

	// 5. Utilizamos *p para desreferenciamento do dado
	// Assim alterando seus dados
	(*p).nome = "Outro nome"
	(*p).idade = 25

}

func main() {
	// 1. Criamos uma pessoa do tipo pessoa
	pessoa1 := pessoa{
		nome:  "Vini",
		idade: 22,
	}

	// 2. Mostramos seu primeiro resultado sem alterações
	fmt.Println(pessoa1)

	// 3. Mandamos o endeço de memória do dado para a função
	mudeMe(&pessoa1)

	// 6. Dados alterados
	fmt.Println(pessoa1)
}
