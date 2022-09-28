package variables

import "fmt"

func Demo1() {
	var metin string = "Bonjour Tout le Monde"
	fmt.Println(metin)
	fmt.Println(metin)

	var tva int = 21
	fmt.Println(tva)
	fmt.Println(100 + (100 * 21 / 100))

	var tva2 float32 = 0.1
	fmt.Println(tva2)
	fmt.Println(100 + 100*tva2)

	tva3 := 25
	fmt.Println(tva3)
	fmt.Printf("Data type : %T\n", tva3)

	var situation bool
	var metin1 string = "Ahmet"
	var metin2 string = "Ahmet"

	situation = metin1 == metin2
	fmt.Println(situation)

}
