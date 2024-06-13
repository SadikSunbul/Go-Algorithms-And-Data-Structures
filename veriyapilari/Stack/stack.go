package main

import "fmt"

func main() {
	/* Dizi ile yapılan Stack
	push(1)
	push(2)
	push(3)
	push(4)
	push(5)
	push(6)
	push(7)

	StackOku()

	fmt.Println()
	fmt.Println("Delete:", pop())
	StackOku()

	fmt.Println()
	fmt.Println("Delete:", pop())
	StackOku()

	fmt.Println()
	fmt.Println("Delete:", pop())
	StackOku() */

	//Bu kısım Baglı listeler kullanilarak yapılmıştır

	stackk := Stackk{veri: 1, önceki: nil}
	Uc = &stackk

	stackk.StackPush(2)
	stackk.StackPush(3)
	stackk.StackPush(4)
	stackk.StackPush(5)
	stackk.StackPush(6)
	stackk.StackPush(7)

	stackk.StackOku()

	fmt.Println()
	fmt.Println("delete:", stackk.StackPop())
	stackk.StackOku()

	fmt.Println()
	fmt.Println("delete:", stackk.StackPop())
	stackk.StackOku()

	fmt.Println()
	fmt.Println("delete:", stackk.StackPop())
	stackk.StackOku()

}

//Stack -1 dizi kulanarak yapılıcak

var Stack []int
var StackSonIndex = -1

func push(deger int) {
	StackSonIndex++
	Stack = append(Stack, deger)
}

func pop() int {
	ındex := StackSonIndex
	StackSonIndex--
	return Stack[ındex]
}

func StackOku() {
	for i := 0; i <= StackSonIndex; i++ {
		fmt.Print(Stack[i], " - ")
	}
}

//Stack -2 bagli liste kullanilcak

type Stackk struct {
	veri   int
	önceki *Stackk
}

var Uc *Stackk

func (s *Stackk) StackPush(veri int) {
	newStack := Stackk{veri: veri, önceki: Uc}
	Uc = &newStack
}

func (s *Stackk) StackPop() int {
	deger := Uc.veri

	Uc = Uc.önceki

	return deger
}

func (s *Stackk) StackOku() {
	tempUc := Uc

	for tempUc != nil {
		fmt.Print(tempUc.veri, " - ")
		tempUc = tempUc.önceki
	}
}
