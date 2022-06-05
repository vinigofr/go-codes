package __estruturas_metodos_intrfaces

import "testing"

type Forma interface {
	Area() float64
}

func TestArea(t *testing.T) {
	type tableTestingStructure []struct {
		forma    Forma
		esperado float64
		nome     string
	}

	// Aprendizados:
	// 1. Podemos fazer testes de tabela e utilizar interfaces e structs para
	//	fazer testes mais diretos ao ponto e com menos linhas de código.
	// 2. É uma boa prática utilizar o nome dos campos quando for montar structs
	// 3. Podemos usar go -run <nome-func-test>/<nome-no-t.Run()>
	//	Ex: go test -run TestArea/Retangulo <- Executará apenas um caso de teste.

	testesArea := tableTestingStructure{
		{forma: Retangulo{Largura: 12, Altura: 6}, esperado: 72, nome: "Retângulo"}, // Primeiro teste
		{Circulo{10}, 314.1592653589793, "Círculo"},                                 // Segundo teste
		{Triangulo{12, 6}, 35, "Triangulo"},                                         // Terceiro teste
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			if resultado != tt.esperado {
				t.Errorf("%#v - resultado %.2f, esperado %.2f", tt.forma, resultado, tt.esperado)
			}
		})
	}
}
