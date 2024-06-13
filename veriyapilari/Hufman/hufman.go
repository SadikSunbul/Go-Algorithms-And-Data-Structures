package main

import (
	"container/heap"
	"fmt"
)

// MinHeapNode, Huffman ağacı düğümlerini temsil eder
type MinHeapNode struct {
	data  rune         // Giriş karakterlerinden biri
	freq  uint         // Karakterin frekansı
	left  *MinHeapNode // Sol çocuk düğüm
	right *MinHeapNode // Sağ çocuk düğüm
}

// CompareMinHeapNode, iki min heap düğümünü karşılaştırmak için kullanılır
type CompareMinHeapNode []*MinHeapNode

func (c CompareMinHeapNode) Len() int           { return len(c) }
func (c CompareMinHeapNode) Less(i, j int) bool { return c[i].freq < c[j].freq }
func (c CompareMinHeapNode) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c *CompareMinHeapNode) Push(x interface{}) {
	*c = append(*c, x.(*MinHeapNode))
}

func (c *CompareMinHeapNode) Pop() interface{} {
	old := *c
	n := len(old)
	item := old[n-1]
	*c = old[0 : n-1]
	return item
}

// printCodes, Huffman ağacının kökünden kodları yazdırır
func printCodes(root *MinHeapNode, str string) {
	if root == nil {
		return
	}

	if root.data != '$' {
		fmt.Println(string(root.data) + ": " + str)
	}

	printCodes(root.left, str+"0")
	printCodes(root.right, str+"1")
}

// HuffmanCodes, Huffman Ağacını oluşturur ve kodları yazdırır
func HuffmanCodes(data []rune, freq []uint, size int) {
	var left, right, top *MinHeapNode

	// MinHeap oluştur ve data[] dizisindeki tüm karakterleri ekler
	minHeap := make(CompareMinHeapNode, size)
	for i := 0; i < size; i++ {
		minHeap[i] = &MinHeapNode{data[i], freq[i], nil, nil}
	}

	// MinHeap'i sırala
	heap.Init(&minHeap)

	// MinHeap'de tek bir düğüm kalmayana kadar döngü yap
	for len(minHeap) != 1 {
		// İki minimum frekanslı düğümü çıkar
		left = heap.Pop(&minHeap).(*MinHeapNode)
		right = heap.Pop(&minHeap).(*MinHeapNode)

		// Frekansları toplamı yeni bir iç düğüm oluştur
		top = &MinHeapNode{'$', left.freq + right.freq, left, right}

		// Yeni düğümü MinHeap'e ekle
		heap.Push(&minHeap, top)
	}

	// Huffman kodlarını yazdır
	printCodes(minHeap[0], "")
}

func main() {
	arr := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	freq := []uint{5, 9, 12, 13, 16, 45}

	size := len(arr)

	HuffmanCodes(arr, freq, size)
}
