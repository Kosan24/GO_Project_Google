package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Urun kaydedildi : ", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Loglandi")
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 5000}

	p = product{productName: "Mouse", unitPrice: 40}
	fmt.Println("Islem basarili")
	fmt.Println(p.productName)
	defer p.save()

}
