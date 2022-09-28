package conditionals

import "fmt"

func Demo2() {
	var compte float64 = 1000
	var demandeArgent float64 = 1000

	if demandeArgent > compte {
		fmt.Println("Fonds insuffisants sur le compte")
	} else if demandeArgent == compte {
		fmt.Println("Reparation votre argent")
		fmt.Println("Attention!! Pas d'argent sur le compte")
	} else {
		fmt.Println("Reparation votre argent")
	}
}
