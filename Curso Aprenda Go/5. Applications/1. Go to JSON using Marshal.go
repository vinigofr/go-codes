package main

import (
	"encoding/json"
	"fmt"
)

// Lembrando que todo struct precisa ter a primeira letra maiúscula quando for
// Necessária a exportação para json
type Pessoa struct {
	Name           string `json:"nome"`
	Sobrenome      string
	Idade          int
	Profissao      string
	Conta_bancaria int
}

func main() {
	vinicius := Pessoa{
		Name:           "Vinicius",
		Sobrenome:      "Gouveia",
		Idade:          22,
		Profissao:      "Dev",
		Conta_bancaria: 1,
	}

	raquel := Pessoa{
		Name:           "Raquel",
		Sobrenome:      "Gouveia",
		Idade:          45,
		Profissao:      "Costureira",
		Conta_bancaria: 40000,
	}

	result_vini, _ := json.Marshal(vinicius)
	result_raquel, _ := json.Marshal(raquel)

	fmt.Println(string(result_vini))
	fmt.Println(string(result_raquel))

}
