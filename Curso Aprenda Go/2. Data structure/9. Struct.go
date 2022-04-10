package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{
		nome:      "Vini",
		sobrenome: "Gouveia",
		fumante:   false,
	}
	fmt.Println(cliente1)
}