# GO DİLİ NOTLARI

Merhaba , oldukça popüler olan Go dilini öğrenmeye başladım ve bir çok eğitim ve araştırma sonucunda öğrendiğim Go dili bilgilerimi buraya aktarıyorum.  Aynı zamanda Medium hesabımda yayınladığım bilgilerime bakabilirsiniz.
Kendimi Go dilinde geliştirdikten sonra çeşitli projelerimde mevcut olucak. 


## İçerik

## Blog Serisi
[Medium](https://medium.com/@didem.kis)


  
## Kullanım/Örnekler

```go
package main //Dosyayı hangi paketin içine koyucağımızı belirleriz.
import "fmt" // Aşağıda ki kodu yazıp Ctrl+ S ile otomatik import ediyor.

func main() { // Go kodlarımızın çalışacağı bölüm
	var metin string = "Go dilinde değişken tanımlamayı öğreniyorum. " //Değişken tanımlamak için var deyimi kullanılır. Hemen sonrasında değişkenin adı ve türü belirtilir.
	fmt.Print(metin)                                                   //Değişkenimizi yazdırıyoruz.
	fmt.Print("Hello ")                                                //Ekrana yazdırma
	fmt.Println("Go dilini öğreniyorum. ")                             //Çıktının alt alta yazdığını görüyoruz .Nedeni ise Println 'de ki ln 'den kaynaklıdır. Line demektir ve bu komut ile alt alta çıktı sağlar.
	fmt.Println("!")

}
