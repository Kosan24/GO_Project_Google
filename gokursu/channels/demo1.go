package channels

import (
	"fmt"
	"time"
)

func NombresPairs(nombrePaireCn chan int) {
	total := 0
	for i := 0; i <= 10; i += 2 {
		total = total + 1
		fmt.Println("Travaillé nombre pair")
		time.Sleep(1 * time.Second)
	}
	nombrePaireCn <- total
}

func NombresimPairs(nombreimPaireCn chan int) {
	total := 0
	for i := 1; i <= 10; i += 2 {
		total = total + 1
		fmt.Println("Travaillé nombre impair")
		time.Sleep(1 * time.Second)
	}
	nombreimPaireCn <- total
}
