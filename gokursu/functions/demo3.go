package functions

func AdditionVariadic(nombres ...int) int {
	addit := 0
	for i := 0; i < len(nombres); i++ {
		addit = addit + nombres[i]
	}

	return addit
}
