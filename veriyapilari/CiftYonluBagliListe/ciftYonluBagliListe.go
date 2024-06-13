package main

import "fmt"

func main() {

	liste := Node{}

	liste.Ekle(5)
	liste.Ekle(6)
	liste.Ekle(75)
	liste.Ekle(58)
	liste.Ekle(59)
	liste.Ekle(35)
	liste.Ekle(51)

	fmt.Println()
	liste.Oku()

	fmt.Println()
	liste.sil(5)
	liste.Oku()

	liste.sil(51)
	fmt.Println()
	liste.Oku()

	fmt.Println()
	liste.sil(59)
	liste.Oku()
}

type Node struct {
	veri    int
	sonraki *Node
	önceki  *Node
}

var Baş *Node
var Son *Node

func (n *Node) Ekle(deger int) {
	_newNode := Node{deger, nil, nil}
	if Baş == nil {
		//bu durumda ılk eleman eklenıyordur
		Baş = &_newNode
		Son = &_newNode
	} else {
		//bu durumda ılk eleman dan farklı bır eleman ekleniyordur
		Son.sonraki = &_newNode
		_newNode.önceki = Son
		Son = Son.sonraki
	}
}

func (n *Node) Oku() {
	if Baş == nil {
		fmt.Println("Liste boş okunacak birşey yok")
		return
	}

	temp := Baş
	for temp != nil {
		fmt.Print(temp.veri, " ")
		temp = temp.sonraki
	}
}

func (n *Node) sil(deger int) {

	if Baş == nil {
		fmt.Println("Ağaç boş")
	}

	if Baş.veri == deger {
		if Son.veri == deger {
			Baş = nil
			Son = nil
			return
		} else {
			Baş = Baş.sonraki
			Baş.önceki = nil
			return
		}
	} else if Son.veri == deger {
		Son = Son.önceki
		Son.sonraki.önceki = nil
		Son.sonraki = nil
	} else {
		temp := Baş
		for temp != nil {
			if temp.veri == deger {
				temp.önceki.sonraki = temp.sonraki
				temp.sonraki.önceki = temp.önceki
				return
			}
			temp = temp.sonraki
		}
	}
}
