package main

func main() {
	x := 10

	// Caso houver mais de um caso correto, serão executadas os respectivos casos

	// Uma das formas pode ser essa
	switch x {
	case 10:
		println("10 sô")
	case 20:
		println("20 sô")
	}

	// Uma outra coisa muito interessante é que Golang te permite ter mais de uma expressão por case
	// A palavra reservada fallthrough permite que o próximo caso seja executado caso sua condição seja verdadeira.

	switch true {
	case true, true:
		println(true)
		fallthrough
	case false, false:
		println(false)
	}

	// O switch também pode ter uma inicialização
	switch e := 10; {
	case e == 10:
		println(true)
	}
}
