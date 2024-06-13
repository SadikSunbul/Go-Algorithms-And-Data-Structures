package main

import (
	"fmt"
)

// encoding, LZW kodlama algoritmasını kullanarak bir string'i kodlar
func encoding(s1 string) []int {
	fmt.Println("Kodlama")

	// Kodlama tablosu oluşturulur
	table := make(map[string]int)
	for i := 0; i <= 255; i++ {
		ch := string(rune(i))
		table[ch] = i
	}

	p, c := "", ""
	p += string(s1[0])
	code := 256
	outputCode := []int{}
	fmt.Println("String\tOutput_Code\tAddition")
	for i := 0; i < len(s1); i++ {
		if i != len(s1)-1 {
			c += string(s1[i+1])
		}
		if _, ok := table[p+c]; ok {
			p += c
		} else {
			fmt.Printf("%s\t%d\t\t%s\t%d\n", p, table[p], p+c, code)
			outputCode = append(outputCode, table[p])
			table[p+c] = code
			code++
			p = c
		}
		c = ""
	}
	fmt.Printf("%s\t%d\n", p, table[p])
	outputCode = append(outputCode, table[p])
	return outputCode
}

// decoding, LZW kodlama algoritmasını kullanarak kodlanmış bir veriyi çözer
func decoding(op []int) {
	fmt.Println("\nÇözme")
	table := make(map[int]string)
	for i := 0; i <= 255; i++ {
		ch := string(rune(i))
		table[i] = ch
	}
	old, n := op[0], 0
	s := table[old]
	c := string(s[0])
	fmt.Print(s)
	count := 256
	for i := 0; i < len(op)-1; i++ {
		n = op[i+1]
		if _, ok := table[n]; !ok {
			s = table[old]
			s += c
		} else {
			s = table[n]
		}
		fmt.Print(s)
		c = string(s[0])
		table[count] = table[old] + c
		count++
		old = n
	}
}

func main() {
	s := "WYS*WYGWYS*WYSWYSG"
	outputCode := encoding(s)
	fmt.Println("Çıktı Kodları: ", outputCode)
	decoding(outputCode)
}
