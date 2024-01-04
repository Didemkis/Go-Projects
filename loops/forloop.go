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
// İKİNCİ KULLANIM

func Ornek2() {
	for i := 1; i < 10; i++ { //ilk başlangıç değerini verdim , ikinci kısım şart kısmı ve i'yi arttırma
		fmt.Println(i)
	}
	fmt.Println("Bitti.")
}

