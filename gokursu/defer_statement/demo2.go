package defer_statement

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		fmt.Println("cift sayi calisti")
		return "Cift sayidir"
	}

	if sayi%2 != 0 {
		return "Tek sayidir"
	}

	return "Belli degil"
}

func Test() {
	sonuc := Demo2(10)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}
