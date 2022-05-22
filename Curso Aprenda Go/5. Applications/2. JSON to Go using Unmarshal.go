package main

import (
	"encoding/json"
	"fmt"
)

// Lembrando que todo struct precisa ter a primeira letra maiúscula quando for
// Necessária a exportação para json
type ViniG struct {
	Name      string `json:"Name"`
	Sobrenome string `json:"Sobrenome"`
}

func main() {
	// Temos aqui o JSON/Objeto

	// Para entender mais, pesquise sobre Tags em Go
	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	vinicius := `{"Name": "vinicius","Sobrenome": "Gouveia"}`

	var animals []string
	var viniParsed ViniG

	json.Unmarshal([]byte(blob), &animals)
	json.Unmarshal([]byte(vinicius), &viniParsed)

	fmt.Println(animals)
	fmt.Println(viniParsed)
}
