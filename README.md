# GO DİLİ NOTLARI

***Merhaba , oldukça popüler olan Go dilini öğrenmeye başladım ve bir çok eğitim ve araştırma sonucunda öğrendiğim Go dili bilgilerimi buraya aktarıyorum.  Aynı zamanda Medium hesabımda yayınladığım bilgilerime bakabilirsiniz.
Kendimi Go dilinde geliştirdikten sonra çeşitli projelerimde mevcut olucak. ***


### İçerik

### Blog Serisi
[Medium](https://medium.com/@didem.kis)

---
  
## Paket Mantığı
- Her kaynak dosya, bu dosyanın hangi pakete ait olduğunu belirten bir paket bildirimi ile başlar. Paket adları her zaman küçük harfle yazılır.
- Bir paketin tüm .go dosyaları tek bir dizinde bulunur. Ayrıca, bir dizinde yalnızca bir paket bulunabilir.
- Main paketi, bir kütüphane yerine bir çalıştırılabilir dosya tanımladığı için özeldir. Her çalıştırılabilir dosya bir main paketi ve giriş noktası olan bir main fonksiyonu içermelidir.
- İçe aktarılan paketler kod bağlandığında eklenir.
- İçe aktarılan paketler yeniden adlandırılabilir (örneğin, import format "fmt", "fmt" paketini "format" olarak yeniden adlandıracaktır).

Main'nimizin içinde diğer dosyalarımızı çağırabilmek için ; 
Öncelikle terminalimizde model oluşturuyoruz , ismine ben golesson dedim başka bir şey konabilir. Main dosyamıza geçip bir sonraki kodu yazarak çağırıyoruz.

```go
go mod init golessson 

package main 
 
import "golesson/Variables" 
func main(){
 Variables.variables //Klasör adımız ve içinde ki dosyanın adı 
}

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