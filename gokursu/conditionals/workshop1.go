package conditionals

import (
	"fmt"
)

func Demo3() {

	var sayi1, sayi2, sayi3 int = 100, 1005, 18

	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}
	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}

	fmt.Printf("Le plus grand nombre : %v", enBuyuk)
}
