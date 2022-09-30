package loops

import "fmt"

func Demo4() {
	numero := 0
	fmt.Println("écrire un nombre")
	fmt.Scanln(&numero)

	nombrePremier := true
	for i := 2; i < numero; i++ {
		if numero%i == 0 {
			nombrePremier = false
		}
	}

	if nombrePremier == true {
		fmt.Println("PremierNombre")
	} else {
		fmt.Println("Pas de Premier Nombre")
	}

}
