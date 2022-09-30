package loops

import "fmt"

func Demo1() {
	var lettre string = "Bonjour toute le monde!!"

	i := 1
	for i <= 5 {
		fmt.Println(lettre)
		i = i + 1
	}
	fmt.Println("Termine")
}
