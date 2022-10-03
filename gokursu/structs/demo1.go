package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Ordinateur", 650, "Dell"})
	fmt.Println(product{name: "Ordinateur", unitPrice: 650, brand: "Dell"})

}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
