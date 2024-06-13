package main

import "fmt"

func main() {

	dizi1 := []int{4, 5, 2, 7, 1, 9}
	fmt.Println(dizi1)
	SelectionSort(dizi1)

	fmt.Println(dizi1)

}

func SelectionSort(dizi []int) { //gerıye dızıyı donmeye gerek yoktur slıce aldıgı ıcın bu fonksıyon referansı gelır buraya kopyası gelmedıgı ıcın gonderılen dızıye erısıp drekt orjnal dızıyı degısrır

	var minIndex int

	for index, _ := range dizi {
		minIndex = index
		for i := index; i < len(dizi); i++ {
			if dizi[i] < dizi[minIndex] {
				minIndex = i
			}
		}
		temp := dizi[index]
		dizi[index] = dizi[minIndex]
		dizi[minIndex] = temp
	}
}
