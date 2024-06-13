package main

import "fmt"

func main() {
	dizi1 := []int{4, 5, 2, 7, 1, 9, 3, 11, 6}
	fmt.Println(dizi1)
	sort(dizi1, 0, len(dizi1)-1)

	fmt.Println(dizi1)
}

func merge(arr1 []int, l, m, r int) {

	n1 := m - l + 1
	n2 := r - m
	L := make([]int, n1)
	R := make([]int, n2)
	var i, j int

	for i = 0; i < n1; i++ {
		L[i] = arr1[l+i]
	}
	for j = 0; j < n2; j++ {
		R[j] = arr1[m+1+j]
	}

	i = 0
	j = 0

	k := l

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr1[k] = L[i]
			i++
		} else {
			arr1[k] = R[j]
			j++
		}
		k++
	}

	for j < n2 {
		arr1[k] = R[j]
		j++
		k++
	}
	for i < n1 {
		arr1[k] = L[i]
		i++
		k++
	}
}

func sort(arr []int, l, r int) {
	if l < r {
		m := l + (r-l)/2

		sort(arr, l, m)
		sort(arr, m+1, r)

		merge(arr, l, m, r)
	}
}
