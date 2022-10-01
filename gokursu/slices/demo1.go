package slices

import "fmt"

func Demo1() {
	lesNoms := make([]string, 3)
	fmt.Println(lesNoms)

	lesNoms[0] = "Ahmet"
	lesNoms[1] = "Hakan"
	lesNoms[2] = "Mehmet"
	lesNoms = append(lesNoms, "Aysel")

	fmt.Println(lesNoms)
	fmt.Println(len(lesNoms))

}
