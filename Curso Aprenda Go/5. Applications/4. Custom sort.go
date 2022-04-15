package main

import (
	"fmt"
	"sort"
)

type Carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordernarPorPotencia []Carro

// Estes três métodos devem existir, pois fazem parte da interface sort.Sort do pacote sort
// Então devemos implementá-lo
func (pot ordernarPorPotencia) Len() int { return len(pot) }
func (pot ordernarPorPotencia) Less(i, j int) bool {
	return pot[i].potencia < pot[j].potencia
}
func (pot ordernarPorPotencia) Swap(i, j int) {
	pot[i], pot[j] = pot[j], pot[i]
}

func main() {
	carros := []Carro{
		{
			nome:     "Hilux",
			potencia: 163,
			consumo:  10,
		},
		{
			nome:     "Gol",
			potencia: 86,
			consumo:  15,
		},
	}

	sort.Sort(ordernarPorPotencia(carros))

	lengthCoisa := ordernarPorPotencia(carros)
	fmt.Println(lengthCoisa, " Numero de carros")

	fmt.Println(carros, " Ordenados por potencia")
}