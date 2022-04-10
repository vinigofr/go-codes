package main

import "fmt"

type pessoa struct {
	Nome      string
	Sobrenome string
	sorvetes  []string
}

func main() {
	pessoa1 := pessoa{
		Nome:      "Vinicius",
		Sobrenome: "Gouveia",
		sorvetes:  []string{"Chocolate", "Maracujá"},
	}

	pessoa2 := pessoa{
		Nome:      "Raquel",
		Sobrenome: "Gouveia",
		sorvetes:  []string{"Morango", "Maracujá"},
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

	for index, value := range pessoa1.sorvetes {
		fmt.Print("indice ", index, " valor ", value, "\n")
	}

	theMap := map[string]pessoa{}

	theMap[pessoa1.Sobrenome] = pessoa1
	theMap[pessoa2.Sobrenome] = pessoa2

	fmt.Println(theMap["Gouveia"])

}