// ŞART BLOKLARI
//Programlamada şart blokları, verilen koşul yerine getirildiğinde bir kod parçası çalıştırılır. Bazen bunlar Kontrol akışı deyimleri olarak da adlandırılır.  Bir programdaki değişikliklere bağlı olarak yürütmenin ilerlemesine ve dallanmasına neden olmak için kullanılır. 
//Go programlamanın karar verme ifadeleri şunlardır:
//if :  Belirli bir şartın veya şart bloğunun yürütülüp yürütülmeyeceğine karar vermek için kullanılır, yani belirli bir koşul doğruysa bir deyim bloğu yürütülür, aksi takdirde yürütülmez.

package conditionals

func Ornek1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen> hesap {
		fmt.Print("Hesaptaki para yetersizdir!")
	}
}