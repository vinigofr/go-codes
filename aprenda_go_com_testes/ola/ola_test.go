package main

// Regras para escrever testes
// Os testes precisam estar em um arquivo chamado xxx_test.go
// A função de teste precisa começar com "Test", func TestGo () {}
// A função de testes recebe um argumento, que é t *testing.T
import "testing"

func TestOla(t *testing.T) {
    resultado := Ola("")
    esperado := "Olá"

    if resultado != esperado {
        t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
    }
}