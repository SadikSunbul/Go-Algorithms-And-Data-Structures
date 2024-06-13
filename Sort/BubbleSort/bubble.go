package main

import "fmt"

func main() {

	dizi1 := []int{4, 5, 2, 7, 1, 9, 3, 11, 6}
	fmt.Println(dizi1)
	bubleSort(dizi1)

	fmt.Println(dizi1)
}

func bubleSort(dizi []int) {

	for i := len(dizi); i >= 0; i-- {
		for j := 0; j < i-1; j++ {
			if dizi[j] > dizi[j+1] {
				dizi[j], dizi[j+1] = dizi[j+1], dizi[j]
			}
		}
	}

}
