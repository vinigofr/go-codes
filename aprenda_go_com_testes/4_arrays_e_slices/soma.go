package __arrays_e_slices

func Soma(numbers []int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func SomaTudo(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		somas = append(somas, Soma(numeros))
	}
	return somas
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int

	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, Soma(final))
		}
	}
	return somas
}
