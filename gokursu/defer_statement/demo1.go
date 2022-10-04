package defer_statement

import "fmt"

func A() {
	fmt.Println("la fonction 'A' a fonctionné")

}
func C() {
	fmt.Println("la fonction 'C' a fonctionné")
}
func D() {
	fmt.Println("la fonction 'D' a fonctionné")
}

func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("la fonction 'B' a fonctionné")

}
