package main

import (
	"golesson/interfaces"
)

func main() {
	//variables.Demo1()
	//fmt.Print()
	//conditionals.Demo3()
	//loops.Demo5()
	//arrays.Demo4()
	//slices.Demo2()
	//functions.DitBonjour("Ahmet")
	//var result = functions.Plus(2, 3)
	//fmt.Println(result * 10)
	// resultat1, resultat2, resultat3, resultat4 := functions.QuatreOperations(10, 6)
	// fmt.Println("Addition : ", resultat1)
	// fmt.Println("Soustraction : ", resultat2)
	// fmt.Println("Multiplication :", resultat3)
	// fmt.Println("Division : ", resultat4)

	//var result = functions.AdditionVariadic(1,4,6,3,10)
	// fmt.Println(functions.AdditionVariadic(1, 4, 6, 3, 10))
	// fmt.Println(functions.AdditionVariadic(1, 4))
	// fmt.Println(functions.AdditionVariadic())

	// nombres := []int{4, 6, 7, 0, 11}
	// fmt.Println(functions.AdditionVariadic(nombres...))
	//maps.Demo1()
	//for_range.Demo3()
	//structs.Demo2()
	// nombrePaireCn := make(chan int)
	// nombreimpaireCn := make(chan int)
	// go channels.NombresPairs(nombrePaireCn)     //cift sayilar
	// go channels.NombresimPairs(nombreimpaireCn) //tek sayilar

	// nombrePaireTotal, nonombreimpaireTotal := <-nombrePaireCn, <-nombreimpaireCn
	// multiplier := nombrePaireTotal * nonombreimpaireTotal
	// fmt.Println("Multiplier : ", multiplier)
	interfaces.Demo2()

}
