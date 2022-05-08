package main

// Regras para escrever testes
// Os testes precisam estar em um arquivo chamado xxx_test.go
// A função de teste precisa começar com "Test", func TestGo () {}
// A função de testes recebe um argumento, que é t *testing.T

/*
   Alguns passos importantes:
    1. Podemos executar uma série de testes através do comando t.Run()
    2. Podemos utilizar funções auxiliares para nos ajudar a diminuir a
        quantidade de linhas nos nossos testes. Um exemplo é a função
        'verificaMensagemCorreta'
*/
import "testing"

func TestOla(t *testing.T) {
    verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
        /*
            t.Helper () é necessário para dizermos ao conjunto de testes que
            este é um método auxiliar. Ao fazer isso, quando o teste falhar, o
            número da linha relatada estará em nossa chamada de função, e não
            dentro do nosso auxiliar de teste. Isso ajudará outros
            desenvolvedores a rastrear os problemas com maior facilidade.
            Se você ainda não entendeu, comente, faça um teste falhar e
            observe a saída do teste.
        */    
        t.Helper()
        if resultado != esperado {
            t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
        }
    }
 
    t.Run("Diz olá para as pessoas", func(t *testing.T) {
        resultado := Ola("Vinicius", "")
        esperado := "Olá, Vinicius"
        verificaMensagemCorreta(t, resultado, esperado)
    });

    t.Run("Diz 'Olá mundo' quando passado uma string vazia", func(t *testing.T) {
        resultado := Ola("", "")
        esperado := "Olá, mundo"
        verificaMensagemCorreta(t, resultado, esperado)
    });

    t.Run("Em espanhol", func(t *testing.T) {
        resultado := Ola("Vinicius", "Espanhol")
        esperado := "Hola, Vinicius"
        verificaMensagemCorreta(t, resultado, esperado)
    })
}