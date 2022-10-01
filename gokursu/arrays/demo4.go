package arrays

import "fmt"

func Demo4() {
	var lesNombres [2][3]int

	lesNombres[0][0] = 1
	lesNombres[0][1] = 3
	lesNombres[0][2] = 5
	lesNombres[1][0] = 2
	lesNombres[1][1] = 4
	lesNombres[1][2] = 6

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(lesNombres[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}

	//fmt.Println(lesNombres[1][1])

}
