package iteracao

func Repetir(caractere string, quantity int) string {
	var quantidadeRepeticoes int = quantity

	repeticoes := ""
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
