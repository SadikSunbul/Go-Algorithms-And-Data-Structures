package main

import "fmt"

//Query baglu liste kullanarak

func main() {

	query := Query{veri: 0, sonraki: nil}

	query.DeQueu(1)
	query.DeQueu(2)
	query.DeQueu(3)
	query.DeQueu(4)
	query.DeQueu(5)

	query.QueryOku()

	fmt.Println()
	fmt.Println("Delete:", query.EnQueu())
	query.QueryOku()

	fmt.Println()
	fmt.Println("Delete:", query.EnQueu())
	query.QueryOku()

	query.DeQueu(6)
	query.DeQueu(7)
	query.DeQueu(8)

	fmt.Println()
	fmt.Println("Delete:", query.EnQueu())
	query.QueryOku()

}

type Query struct {
	veri    int
	sonraki *Query
}

var Baş *Query
var Son *Query

func (q *Query) DeQueu(deger int) { //ekleme işlemi yapıcak
	newQuery := Query{veri: deger, sonraki: nil}
	if Baş == nil {
		Baş = &newQuery
		Son = &newQuery
	} else {
		Son.sonraki = &newQuery
		Son = &newQuery
	}
}

func (q *Query) QueryOku() {

	tempBas := Baş

	for tempBas != nil {
		fmt.Print(tempBas.veri, " - ")
		tempBas = tempBas.sonraki
	}
}

func (q *Query) EnQueu() int { //cıkarma işlemi yapıcak

	//Bastan okunmaya baslar ve sılme ıslemı oradan yapılır
	deger := Baş.veri

	Baş = Baş.sonraki

	return deger
}
