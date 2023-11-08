package loops

import "fmt"

func Ornek1() {
	var metin string = "Merhaba dünya"

	i := 1
	//infinite loop
	for i <= 5 {
		fmt.Println(metin)
		i = i + 1
	}
	fmt.Println("bitti")
}
