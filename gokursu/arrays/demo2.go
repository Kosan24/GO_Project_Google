package arrays

import "fmt"

func Demo2() {
	var villes [5]string
	villes[0] = "Bruxelles"
	villes[1] = "Antwerpen"
	villes[2] = "Liege"
	villes[3] = "Mons"
	villes[4] = "Verviers"

	fmt.Println(villes)

	for i := 0; i < 5; i++ {
		fmt.Println(villes[i])
	}
}
