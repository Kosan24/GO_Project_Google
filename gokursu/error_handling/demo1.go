package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo11.txt")
	//nil
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadi : ", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadi")
			return
		}

	} else {
		fmt.Println(f.Name())
	}
}
