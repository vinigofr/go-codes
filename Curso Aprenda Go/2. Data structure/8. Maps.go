package main

import "fmt"

func main() {

	// Key: value, looks like json
	contacts := map[string]int{
		"Raquel":   1,
		"Vinicius": 2,
	}

	fmt.Println(contacts)

	// Pode ser acessado via hash
	fmt.Println(contacts["Raquel"])

	// Chaves novas podem ser criada
	contacts["Andréia"] = 3
	fmt.Println(contacts["Andréia"])

	// Caso um valor inexistente no map seja acessado, retornará zero
	fmt.Println(contacts["Desconhecido"])

	// Ok, mas as vezes precisaremos realmente ter o valor zero, o que fazer?

	// Uma alternativa é poder desestruturar os valores de um índice.
	// O primeiro valor é o próprio valor e o segundo é um status
	// Esse status tem como valor um booleano.
	// Caso false, significa que não há nada na posição alvo que foi reauisitada
	// Caso true, a posição já está em uso

	// Isso é chamado de "comma ok idiom"
	value, ok := contacts["Desconhecido"]
	fmt.Println(value) // valor
	fmt.Println(ok) // boolean

	// Percorrendo sobre um map
	for key, value := range contacts {
		fmt.Println(key, value)
	}

	// Podemos também remover um item de um mapa:
	delete(contacts, "Vinicius")
	fmt.Println(contacts)
	// O elemento com a chave "Vinicius" não será exibido
}