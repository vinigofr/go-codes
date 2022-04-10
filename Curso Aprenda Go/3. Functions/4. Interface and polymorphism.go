// Aparemente, interfaces são similares à classes em outras linguagens.
// Uma interface é um conjunto de métodos, ou seja, classe.

// Uma interface permite que um valor possua mais que um tipo
package main

import "fmt"

// 1. Primeiro montamos o tipo base pessoa
type pessoa struct {
	nome string
	sobrenome string
	idade int
}

// 2. Fazemos com que dentista e arquiteto herdem as propriedades de pessoa
type dentista struct {
	pessoa
	atendimentos_realizados string
	salario float64
}

type arquiteto struct {
	pessoa
	salario float64
	tipo_build string
}

// 3. Montamos o método oibomdia() para ambos casos. Observe o "receiver" da função oibomdia()
func (p dentista) oibomdia() {
	fmt.Println("Bom dia do dentista")
}

func (p arquiteto) oibomdia() {
	fmt.Println("Bom dia do arquiteto")
}

func main() {
	// 4. Montamos a estrutura com os dados específicos com cada especialização da pessoa
	umDentista := dentista{
		pessoa: pessoa{
			nome: "Vinicius",
			sobrenome: "Gouveia",
			idade: 22,
		},
		atendimentos_realizados: "50",
		salario: 105,
	}

	umArquiteto := arquiteto{
		pessoa: pessoa{
			nome: "Vinicius",
			sobrenome: "Gouveia",
			idade: 22,
		},
		salario: 105,
		tipo_build: "grandes prédios",
	}

	// Arui podemos ter acesso ao método oibomdia() de cada pessoa.
	umDentista.oibomdia()
	umArquiteto.oibomdia()

	// 3.1 Quando precisamos chamar serHumano(), podemos chamar mandar a estrutura de dado 'dentista' ou "arquiteto"
	serHumano(umArquiteto)
	// 4. Pelo fato de que a função lá dentro se chama oibomdia(), ela será executada.
	serHumano(umDentista)

}


// ====== Um outro caso que pode ser feito, é a criação de interface.
// 1. Criamos um tipo "gente" e montamos quais métodos deverão ser usados.
type gente interface {
	oibomdia()
}

// 2. Criamos uma função que recebe coisas do tipo "gente"
// 3. Olha para a função main.
func serHumano(g gente) {
	g.oibomdia()

	// Podemos também ter regras condicionais com Switch case para verificação de tipos específicos e realizar ações a partir disso.

	switch g.(type) {
	case dentista:
		fmt.Println("Deu dentista");
	case arquiteto:
		fmt.Println("Deu arquiteto");
	}
}




