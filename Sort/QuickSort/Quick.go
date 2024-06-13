package main

import "fmt"

func main() {
	dizi1 := []int{4, 5, 2, 7, 1, 9, 3, 11, 6}
	fmt.Println(dizi1)
	quicSort(dizi1, 0, len(dizi1)-1)

	fmt.Println(dizi1)
}

func partion(arr []int, low, hight int) int {

	pivot := arr[hight]

	i := low - 1

	for j := low; j <= hight-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[hight] = arr[hight], arr[i+1]
	return i + 1
}

func quicSort(arr []int, low, high int) {

	if low < high {
		pi := partion(arr, low, high)

		quicSort(arr, low, pi-1)
		quicSort(arr, pi+1, high)
	}
}
