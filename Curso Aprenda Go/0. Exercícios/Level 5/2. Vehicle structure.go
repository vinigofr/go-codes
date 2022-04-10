package main

import "fmt"

type veículo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veículo
	quatroRodas bool
}

type sedan struct {
	veículo
	modeloLuxo bool
}

func main() {
	hilux := caminhonete{
		quatroRodas: true,
		veículo: veículo{
			portas: 4,
			cor:    "branca",
		},
	}

	siena := sedan{
		veículo: veículo{
			portas: 4,
			cor:    "branca",
		},
		modeloLuxo: false,
	}

	fmt.Println(hilux)
	fmt.Println(siena)

	
}