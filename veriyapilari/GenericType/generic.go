package main

import "fmt"

func main() {
	// Int tipinde bir slice tanımlayalım.
	intSlice := []int{1, 2, 3, 4, 5}
	// String tipinde bir slice tanımlayalım.
	stringSlice := []string{"apple", "banana", "orange"}

	// Generic fonksiyonumuzu çağıralım ve sonucu ekrana basalım.
	fmt.Println(ReverseSlice(intSlice))    // [5 4 3 2 1]
	fmt.Println(ReverseSlice(stringSlice)) // [orange banana apple]
}
func ReverseSlice[T any](s []T) []T {
	// Tersine çevrilecek slice'in uzunluğunu alalım.
	length := len(s)
	// Yeni bir slice oluşturalım ve kapasitesini orijinal slice'in uzunluğuyla aynı olarak belirleyelim.
	reversed := make([]T, length, length)
	// Orijinal slice'i tersine çevirerek yeni slice'e kopyalayalım.
	for i, v := range s {
		reversed[length-i-1] = v
	}
	return reversed
}
