package main

const vinicius = "Vinicius"

var vinicius2 = 'v'

func main() {
	switch esporteFavorito := "Correr"; {
	case esporteFavorito == "Correr":
		println("Vamos correr")
	default:
		println("Não vamos praticar nenhum esporte!")
	}
}