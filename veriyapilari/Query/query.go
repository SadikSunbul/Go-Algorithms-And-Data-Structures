package query

import "fmt"

//Query baglu liste kullanarak

func main() {

	query := Query{veri: 0, sonraki: nil}

	query.EnQueu(1)
	query.EnQueu(2)
	query.EnQueu(3)
	query.EnQueu(4)
	query.EnQueu(5)

	query.QueryOku()

	fmt.Println()
	fmt.Println("Delete:", query.DeQueu())
	query.QueryOku()

	fmt.Println()
	fmt.Println("Delete:", query.DeQueu())
	query.QueryOku()

	query.EnQueu(6)
	query.EnQueu(7)
	query.EnQueu(8)

	fmt.Println()
	fmt.Println("Delete:", query.DeQueu())
	query.QueryOku()

}

type Query struct {
	veri    interface{}
	sonraki *Query
}

var Baş *Query
var Son *Query
var count int = 0

func (q *Query) EnQueu(deger interface{}) { //ekleme işlemi yapıcak
	newQuery := Query{veri: deger, sonraki: nil}
	if Baş == nil {
		Baş = &newQuery
		Son = &newQuery
	} else {
		Son.sonraki = &newQuery
		Son = &newQuery
	}
	count++
}

func (q *Query) QueryOku() {

	tempBas := Baş

	for tempBas != nil {
		fmt.Print(tempBas.veri, " - ")
		tempBas = tempBas.sonraki
	}
}

func (q *Query) DeQueu() interface{} { //cıkarma işlemi yapıcak

	//Bastan okunmaya baslar ve sılme ıslemı oradan yapılır
	deger := Baş.veri

	Baş = Baş.sonraki
	count--
	return deger
}

func (q *Query) Count() int {
	return count
}
