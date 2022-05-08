package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {
    soma := Adiciona(2, 2)
    esperado := 4

    if soma != esperado {
        t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
    }
}


// Regras e informacões para execução de uma função exemplo:
/*
    1. Também precisam estar em um arquivo *_test.go
    2. Podem ser executadas diretamente da documentação
    3. As funções de exemplo obrigatoriamente devem começar com o prefixo Example
*/

func ExampleAdiciona() {
    fmt.Println(Adiciona(1, 1))
    // Output: 2
}