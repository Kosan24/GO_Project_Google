package arrays

import "fmt"

func Demo3() {
	numeros1 := [5]int{20, 30, 50, 10, 2}
	fmt.Println(numeros1)

	plusGrande := numeros1[0]

	//lenght = uzunluk
	for i := 0; i < len(numeros1); i++ {
		if numeros1[i] > plusGrande {
			plusGrande = numeros1[i]
		}
		fmt.Println("Plus Grande : ", plusGrande)
	}

}
