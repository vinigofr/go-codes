// Sim, como soube?
package main

import "fmt"

// Oi, teste
func Ola(nome, idioma string) string{
    var prefixoPT = "Olá, "
    var prefixoES = "Hola, "
    if nome == "" {
        nome = "mundo"
    }

    if idioma == "Espanhol" {
        return prefixoES + nome
    }

    return prefixoPT + nome
}

func main() {
    fmt.Println(Ola("", "espanhol"))
}