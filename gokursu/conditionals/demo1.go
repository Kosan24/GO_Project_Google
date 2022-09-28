package conditionals

import "fmt"

func Demo1() {
	var compte float64 = 1000
	var demandeArgent float64 = 900

	if demandeArgent > compte {
		fmt.Println("Fonds insuffisants sur le compte")
	}

	if demandeArgent <= compte {
		fmt.Println("Reparation votre argent")
		compte = compte - demandeArgent
	}
	fmt.Printf("Terminer. Argent en compte : %v", compte)

}
