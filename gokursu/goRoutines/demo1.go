package goroutines

import (
	"fmt"
	"time"
)

func NombresPairs() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Nombre pairs :", i) //cift sayilar
		time.Sleep(1 * time.Second)
	}
}

func NombresimPairs() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Nombre impairs :", i) //tek sayilar
		time.Sleep(1 * time.Second)
	}
}
