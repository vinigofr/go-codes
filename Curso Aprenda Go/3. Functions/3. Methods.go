package main

import "fmt"

// Métodos são funções anexadas a um tipo
// Quando anexamos uma função a um tipo, ela se torna um método daquele tipo

type pessoa struct {
	nome string
	idade int
}

// Usando receivers em funções, podemos especificar o que uma função pode acabar recebendo.
// O receiver demonstra o contexto da função
func (p pessoa) funcaoPessoa() (int, string) {
	return p.idade, p.nome
}

func main() {
	// Primeiro você instancia uma pessoa do tipo pessoa e guarda seus dados
	pessoa1 := pessoa{"Vinicius", 22}
	fmt.Println(pessoa1)

	// Segundo, você consegue utilizar a função funcaoPessoa() que foi atrelada ao tipo pessoa.
	// Assim, você pode desestruturar os dados que foram passados logo acima na linha 21
	idade, nome := pessoa1.funcaoPessoa()
	fmt.Println(idade)
	fmt.Println(nome)

	// Chamar a funcão em outro lugar não funcionará, POIS ELE É UMA FUNÇÃO ESPECÍFICA PARA VALORES DAQUELE TIPO
}