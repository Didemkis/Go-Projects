// Programlama da bir döngüyü, bir kod bloğunu tekrarlamak için kullanırız. Mesela, bir ifadeyi 100 kez yazdırmak istiyorsak, aynı print ifadesini 100 kez yazmak yerine, aynı kodu 100 kez çalıştırmak için bir döngü kullanabiliriz.
// For döngülerini karmaşık programları temizlemek ve basitleştirmek için kullanırız.
// Döngü başladığında, başlangıç değeri belirtilen değerle başlar. Ardından, koşul değerlendirilir. Koşul doğru (true) ise, döngü içindeki işlemler çalıştırılır ve ardından artış/değişiklik adımı gerçekleşir. Koşul hala doğruysa, döngü işlemi tekrarlanır. Koşul yanlış (false) olduğunda, döngü sona erer ve program akışı döngüden sonraki kod satırına geçer.
package loops

import "fmt"

func Ornek1() {
	var metin string = "Merhaba dünya"

	i := 1
	//infinite loop : Sonsuz bir döngü, döngü koşulunun sürekli olarak "true" olarak değerlendirilmesi nedeniyle asla sona ermeyen bir döngüdür.Sonsuz döngüyü sonlandırmak için programı manuel olarak sonlandırmanız veya bir kesme (break) ifadesi veya başka bir kontrol mekanizması eklemeniz gerekecektir.
	for i <= 5 {
		fmt.Println(metin)
		i = i + 1
	}
	fmt.Println("bitti")
}
