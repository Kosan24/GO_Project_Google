package for_range

import "fmt"

//1-10 arasindaki sayilardan tek olanlari toplayan program
//for-range
func Demo2() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total := 0
	for _, numero := range numeros {
		if numero%2 != 0 {
			total = total + numero
		}
	}
	fmt.Println("Total : ", total)
}
