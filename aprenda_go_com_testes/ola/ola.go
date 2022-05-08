// Sim, como soube?
package main

import "fmt"

// Oi, teste
func Ola(value string) string{
    if value != "" {
        return fmt.Sprintf("Olá %v", value)
    }
    return "Olá"
}

func main() {
    fmt.Println(Ola(""))
}