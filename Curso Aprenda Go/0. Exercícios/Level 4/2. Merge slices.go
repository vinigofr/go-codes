package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x[:3], x[6:]...)
	fmt.Println(x)


	states := make([]string, 27, 27)
	states = append(states, "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins")


	fmt.Print("Estados", states, "\n")
	fmt.Print("Capacidade", cap(states), "\n")
	fmt.Print("Tamanho", len(states), "\n")

}