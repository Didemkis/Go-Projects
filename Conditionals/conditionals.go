// ŞART BLOKLARI
//Programlamada şart blokları, verilen koşul yerine getirildiğinde bir kod parçası çalıştırılır. Bazen bunlar Kontrol akışı deyimleri olarak da adlandırılır.  Bir programdaki değişikliklere bağlı olarak yürütmenin ilerlemesine ve dallanmasına neden olmak için kullanılır.
//Go programlamanın karar verme ifadeleri şunlardır:
//if :  Belirli bir şartın veya şart bloğunun yürütülüp yürütülmeyeceğine karar vermek için kullanılır, yani belirli bir koşul doğruysa bir deyim bloğu yürütülür, aksi takdirde yürütülmez.
//Yalnızca if deyimi bize koşul doğruysa ifade bloğunu çalıştıracağını söyler, ancak koşul yanlışsa bunu yapmaz. Peki ya koşul yanlışsa ve  biz başka bir işlem yapmak istiyorsak? İşte else deyimi burada devreye giriyor. Belirtilen koşulumuz yanlış olsada başka bir kod bloğunu çalıştırmanıza olanak tanır.

package Conditionals

import "fmt"

//Basit bir not sisteminden örnek vereyim eğer puanı 50'nin altındaysa geçemedi.
//Eşit veya yüksek ise geçtiğini ekrana çıktılar.
func Ornek1() {
	var gecernot float64 = 50
	var alınan float64 = 48

	if gecernot > alınan {
		fmt.Println("Geçemediniz!")
	}
	if gecernot <= alınan {
		fmt.Println("Geçtiniz!")

	}
	fmt.Println("Bitti.")
}

//Banka hesabımızdan para çekmek istiyoruz ve ona göre bize sonuç veriyor.
func Ornek2() {
	var hesap float64 = 1000
	var cekılmekIstenen float64 = 800

	if cekılmekIstenen > hesap {
		fmt.Println("Hesabınızda yeterli para yoktur!")

	}

	if cekılmekIstenen <= hesap {
		fmt.Println("Paranız hazırlanıyor.")
		hesap = hesap - cekılmekIstenen
	}
	fmt.Println("Bitti. Hesabınızda ki miktar : " + fmt.Sprintf("%v", hesap))
	//Sprintf , değişkenlere formatlı atama yapmamıza yardımcı olur.
}

func Ornek3() {
	var hesap float64 = 1000
	var cekılmekIstenen float64 = 1000

	if cekılmekIstenen > hesap {
		fmt.Println("Hesabınızda yeterli para yoktur!")
	} else if cekılmekIstenen == hesap {
		fmt.Println("Paranız hazırlanıyor.")
		fmt.Println("Hesabınızda para kalmadı!!")
	} else {
		fmt.Println("Paranız hazırlanıyor.")
	}
}
