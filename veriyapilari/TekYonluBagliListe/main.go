package main

import "fmt"

//Tek yonlu baglı lıste yazılıcak

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
	deger   int
	sonraki *Node
}

var Baş *Node
var Son *Node

func (n *Node) Ekle(deger int) {

	_newNode := Node{deger, nil}
	if Baş == nil {
		//bu durumda ılk eleman eklenıyordur
		Baş = &_newNode
		Son = &_newNode
	} else {
		//bu durumda ılk eleman dan farklı bır eleman ekleniyordur
		Son.sonraki = &_newNode
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
		fmt.Print(temp.deger, " ")
		temp = temp.sonraki
	}
}

func (n *Node) sil(deger int) {

	if Baş == nil {
		fmt.Println("Ağaç boş")
	}

	if Baş.deger == deger {
		if Son.deger == deger {
			Baş = nil
			Son = nil
			return
		} else {
			Baş = Baş.sonraki
			return
		}
	} else if Son.deger == deger {
		temp := Baş
		for temp != nil {
			if temp.sonraki.deger == deger {
				temp.sonraki = nil
				Son = temp
				return
			}
			temp = temp.sonraki
		}
	} else {
		temp := Baş
		for temp != nil {
			if temp.sonraki.deger == deger {
				temp.sonraki = temp.sonraki.sonraki
				return
			}
			temp = temp.sonraki
		}
	}
}
