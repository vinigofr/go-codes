// Sim, como soube?
package main

import "fmt"

// Oi, teste
func Ola(nome, idioma string) string{
    if nome == "" {
        nome = "mundo"
    }

    return prefixoSaudacao(idioma) + nome    
}

func prefixoSaudacao(idioma string) (prefixo string){
    const (
        prefixoPT = "Olá, "
        prefixoES = "Hola, "
        prefixoFR = "Bonjour, "
    )
    
    var prefixoAtual string;

    switch idioma {
    case "Francês":
        prefixoAtual = prefixoFR
    case "Espanhol":
        prefixoAtual = prefixoES
    default:
        prefixoAtual = prefixoPT
    }

    return prefixoAtual
}

func main() {
    fmt.Println(Ola("", "espanhol"))
}