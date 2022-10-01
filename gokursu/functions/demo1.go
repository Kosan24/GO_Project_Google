package functions

import "fmt"

func Plus(numero1 int, numero2 int) int {
	var pluss = numero1 + numero2
	return pluss
}

func DitBonjour(nomdUtilisateur string) {
	fmt.Println("Bonjour ", nomdUtilisateur)
}
