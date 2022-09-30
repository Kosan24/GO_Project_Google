package loops

import "fmt"

//220 ve 284 arkadas sayilardir.
//10 ve 65 arkadas sayi degildir.
func Demo5() {
	nombre1 := 220
	nombre2 := 284
	total1 := 0
	total2 := 0

	for i := 1; i < nombre1; i++ {
		if nombre1%i == 0 {
			total1 = total1 + i
		}
	}

	for i := 1; i < nombre2; i++ {
		if nombre2%i == 0 {
			total2 = total2 + i
		}
	}

	if total1 == nombre2 && total2 == nombre1 {
		fmt.Printf("%v et %v les amis sont des nombres", nombre1, nombre2)
	} else {
		fmt.Printf("%v et %v les amis ne sont pas des numéros", nombre1, nombre2)
	}

}
