package main

import "fmt"

func main() {
	dizi1 := []int{4, 5, 2, 7, 1, 9, 3, 11, 6}
	fmt.Println(dizi1)
	sort(dizi1)

	fmt.Println(dizi1)
}

func sort(arr []int) {

	N := len(arr)
	for i := N/2 - 1; i >= 0; i-- {
		heapify(arr, N, i)
	}
	for i := N - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, N, i int) {
	larges := i
	l := 2*i + 1
	r := 2*i + 2

	if l < N && arr[l] < arr[larges] {
		larges = l
	}

	if r < N && arr[r] < arr[larges] {
		larges = r
	}

	if larges != i {
		arr[i], arr[larges] = arr[larges], arr[i]
		heapify(arr, N, larges)
	}
}
