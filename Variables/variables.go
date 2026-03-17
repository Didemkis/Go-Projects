package variables

import "fmt" // Aşağıda ki kodu yazıp Ctrl+ S ile otomatik import ediyor.

func variables(){// Go kodlarımızın çalışacağı bölüm

fmt.Print("Hello ")                    //Ekrana yazdırma
fmt.Println("Go dilini öğreniyorum. ") //Çıktının alt alta yazdığını görüyoruz .Nedeni ise Println 'de ki ln 'den kaynaklıdır. Line demektir ve bu komut ile alt alta çıktı sağlar.
fmt.Println("!")

//DEĞİŞKEN TANIMLAMA
var metin string = "Go dilinde değişken tanımlamayı öğreniyorum. " //Değişken tanımlamak için var deyimi kullanılır. Hemen sonrasında değişkenin adı ve türü belirtilir. Açık değişken atama işlemi.
fmt.Println(metin)                                                 //Değişkenimizi yazdırıyoruz.

hesap := 100 // Kapalı değişken atama.
fmt.Println(hesap)
// Veri tipi öğrenmek istersek
fmt.Printf("veri tipi : %T", hesap) //
// Printf = formatını yani veri tipini bize söyler ."%T"  yüzde işareti kendinden sonraki ilkini yazdırır T işareti de type'ını yazdırır

 // ":=" değişkeni hem tanımlıyorum hemde atıyorum demek. 
 
//Örnek1
var durum bool

var metin1 string = "didem"
var metin2 string = "karamel"

durum = metin1 == metin2
fmt.Println(durum)
}