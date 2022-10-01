package for_range

import "fmt"

func Demo1() {
	villes := []string{"Bruxelles", "Namur", "Liege"}
	for i := 0; i < len(villes); i++ {
		fmt.Println(villes[i])
	}

	for _, ville := range villes {
		fmt.Println(ville)
	}

	for i, ville := range villes {
		fmt.Print(i)
		fmt.Println(ville)
	}
}
