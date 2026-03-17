package loops

import "fmt"

//Ufak bir tahmin oyunu örneği yapıyorum.
func Demo1() {
	TuttugumSayi := 25 //Öncelikle bir değişten atması yapıyorum. Amaç bu sayıyı bilmek.
	Tahmin := 0

	//Kullanıcıdan tahmini alabilmek için scanln fonksiyonunu kulanıyoruz.
	fmt.Println("Lütfen bir sayı tahmin ediniz!")
	fmt.Scanln(&Tahmin) //Scanln : Birden fazla değer okumak için kullanılır.

	for TuttugumSayi != Tahmin { //Tahmin edilen sayı olana kadar kod bloğunu çalıştır.
		if TuttugumSayi < Tahmin {
			fmt.Println("Daha büyük sayı tahmin ediniz!")
			fmt.Scanln(&Tahmin)
		}
		if TuttugumSayi > Tahmin {
			fmt.Println("Daha küçük sayı tahmin ediniz!")
			fmt.Scanln(&Tahmin)
		}

	}
	fmt.Println("Tebrikler bildiniz.")

}

//Kullanıcıdan bir sayı girmesini iste ve girilen sayının asal sayı olup olmadığını belirt.
func Demo2() {
	//1. aşama
	sayi := 0
	fmt.Println("Lütfen bir sayı giriniz.")
	fmt.Scanln(&sayi)

	durum := true
	for i := 2; i < sayi; i++ { //2 'den kendisine kadar olan sayılara bölme.
		if sayi%i == 0 { // % mod operatörü. Sayının i'ye bölümünden kalan 0 ise 
			durum = false
		}
	}

	if durum == true {
		fmt.Println("Asaldır.")
	} else {
		fmt.Println("Asal değildir.")
	}
}
