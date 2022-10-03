package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("Ajoutée : ", c.firstName)
}
func (c customer) update() {
	fmt.Println("Actualisé : ", c.firstName)
}

func Demo2() {
	c := customer{firstName: "Ahmet", lastName: "KOSAN", age: 38}
	c.save()
	c.update()
}
