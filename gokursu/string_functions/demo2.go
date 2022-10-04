package string_functions

import (
	"fmt"
	s "strings" // s harfi allias yani yerine olarak kullanilir
)

func Demo2() {
	prenom := "Ahmet"
	fmt.Println(s.HasPrefix(prenom, "Ahm"))
	fmt.Println(s.HasSuffix(prenom, "et"))
	fmt.Println(s.Index(prenom, "me"))

	desLettres := []string{"a", "h", "m", "e", "t"}
	fmt.Println(s.Join(desLettres, "*"))
	egale := s.Join(desLettres, "*")
	fmt.Println(egale)

	fmt.Println(s.Replace(egale, "*", "+", 3))
	fmt.Println(s.Split(egale, "-"))
	fmt.Println(s.Repeat(egale, 5))

}
