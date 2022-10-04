package string_functions

import (
	"fmt"
	s "strings" // s harfi allias olarak kullanilir
)

func Demo1() {
	prenom := "Ahmet"
	fmt.Println(s.Count(prenom, "h"))
	fmt.Println(s.Contains(prenom, "t"))
	egale := s.Index(prenom, "g")

	if egale != -1 {
		fmt.Println("il y a la lettre g")
	} else {
		fmt.Println("Aucun la lettre g")
	}

	fmt.Println(s.ToLower(prenom))
	fmt.Println(s.ToUpper(prenom))

}
