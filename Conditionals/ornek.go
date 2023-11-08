package Conditionals

import "fmt"

// En büyük sayıyı bulma
func Demo1() {
	var sayi1, sayi2, sayi3 int = 100, 101, 98
	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2

	}
	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}
	fmt.Printf("en büyük sayı : %v", enBuyuk)

}
