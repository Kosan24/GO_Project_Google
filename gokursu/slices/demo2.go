package slices

import "fmt"

func Demo2() {
	lesVille := []string{"Bruxelles", "Liege", "Namur"}
	fmt.Println(lesVille)
	lesVillecopy := make([]string, len(lesVille))
	fmt.Println(lesVillecopy)
	copy(lesVillecopy, lesVille)
	fmt.Println(lesVillecopy)
	lesVille = append(lesVille, "Mons")
	fmt.Println(lesVille)
	fmt.Println(lesVillecopy)

	fmt.Println(lesVille[1:3])
	fmt.Println(lesVille[1:])
	fmt.Println(lesVille[:2])

}
