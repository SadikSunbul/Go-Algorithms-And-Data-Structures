package main

import "fmt"

func main() {
	dizi1 := []int{4, 5, 2, 7, 1, 9, 3, 11, 6}
	fmt.Println(dizi1)
	InsertionSort(dizi1)

	fmt.Println(dizi1)
}

func InsertionSort(dizi []int) {

	for i := 1; i < len(dizi); i++ {
		for j := i; j > 0 && dizi[j] < dizi[j-1]; j-- {
			if dizi[j-1] > dizi[j] {
				dizi[j-1], dizi[j] = dizi[j], dizi[j-1]
			}
		}
	}
}
