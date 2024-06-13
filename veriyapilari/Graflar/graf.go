package main

import "fmt"

//Burda dizi kullanılarak yapıldı

func main() {
	GrafıBaşlat(5)
	BağlantıOluştur(1, 2)
	BağlantıOluştur(2, 2)
	BağlantıOluştur(3, 2)
	BağlantıOluştur(3, 3)
	BağlantıOluştur(3, 4)
	BağlantıKaldır(3, 4)
	fmt.Print()
}

var Graf [][]int

/* Komsuluk matrısı ıle yapıcaz işlemlerimizi
	  1	 2  3  4
    1
    2
	3
	4
*/

func GrafıBaşlat(dugumSayisi int) {
	Graf = make([][]int, dugumSayisi)
	for i := range Graf {
		Graf[i] = make([]int, dugumSayisi)
	}
	//GrafıBaşlat fonksiyonunda Graf değişkenine boyut atanıyor,
	//ancak her bir satırın boyutu atanmıyor. GrafıBaşlat fonksiyonunda,
	//her bir satır için slice oluşturmanız gerekir.
}

func BağlantıOluştur(aKonumundan, bKonumuna int) {
	Graf[aKonumundan][bKonumuna] = 1
}

func BağlantıKaldır(aKonumundan, bKonumuna int) {
	Graf[aKonumundan][bKonumuna] = 0
}
