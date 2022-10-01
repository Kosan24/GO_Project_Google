package for_range

import "fmt"

func Demo3() {
	dictionnaire := map[string]string{"book": "livre", "car": "voiture"}
	for k, v := range dictionnaire {
		fmt.Println(k) //key leri verir
		fmt.Println(v) //degerleri verir
	}
}
