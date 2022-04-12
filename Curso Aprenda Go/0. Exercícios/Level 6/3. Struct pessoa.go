package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type gente interface {
	mostrarPessoa()
}

func helloPerson(g gente) {
	g.mostrarPessoa();
}

func (p pessoa) mostrarPessoa() {
	fmt.Printf("Oi, meu nome é %v %v e tenho %v anos", p.nome, p.sobrenome, p.idade);
}

func main() {
	pessoa1 := pessoa{
		nome:      "Vinicius",
		sobrenome: "Gouveia",
		idade:     22,
	}

 	pessoa1.mostrarPessoa()
	helloPerson(pessoa1)
}