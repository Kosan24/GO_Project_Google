package maps

import "fmt"

func Demo1() {

	dictionnaire := make(map[string]string)
	dictionnaire["table"] = "Table"
	dictionnaire["book"] = "Livre"
	dictionnaire["pencil"] = "Bic"

	fmt.Println(dictionnaire["book"])
	fmt.Println("Nombre d'éléments : ", len(dictionnaire))
	fmt.Println("Dictionnaire : ", dictionnaire)
	delete(dictionnaire, "book")
	fmt.Println("Nombre d'éléments : ", len(dictionnaire))
	fmt.Println("Dictionnaire : ", dictionnaire)

	value, ilya := dictionnaire["table"]
	fmt.Println(value)
	fmt.Println("Statut d'être sur la liste : ", ilya)

	dictionnaire2 := map[string]string{"glass": "verre", "bed": "lit"}
	fmt.Println(dictionnaire2)

}
