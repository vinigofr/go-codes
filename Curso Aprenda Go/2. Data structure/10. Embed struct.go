package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome:  "Vini",
		idade: 22,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Orlando",
			idade: 18,
		},
		titulo:  "Programador",
		salario: 21000,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}