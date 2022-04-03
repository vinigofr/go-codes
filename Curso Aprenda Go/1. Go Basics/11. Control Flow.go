package main

import "fmt"

// Uma maneira de repetir instruções é usando um for loop.
// Não existe while no Go, infelizmente ;-;
// O ; (PONTO E VÍRGULA) em GO é adicionado automaticamente quando compilado.
// É como se fosse em C, mas em go não é obrigatório.

func main() {
    // Exemplos de while

    // 1. Declaração de uma var auxiliar
    y := 0

    // While true + condição de parada
    for y < 10 {
        // Adição
        y++
    }

    // 2. While true forever
    for {
        // A única coisa que pode parar este for é um break
        break
    }

    // 3. For loop simples
	for x := 0; x < 10; x++ {
		fmt.Printf("Hello\n");
	}

    // 4. Ignore uma condição com o continue
    for i := 0; i < 20; i++ {
        if i % 2 != 0 {
            continue
        }
        fmt.Println(i)
    }
}