package main

import "fmt"

func main() {
	const MAX = 100
	var pemenang [MAX]string
	var nomor [MAX]string
	var jumlah, skorA, skorB, i int
	var klubA, klubB string
	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)
	i = 1
	for {
		fmt.Print("Pertandingan ", i, " : ")
		fmt.Scan(&skorA, &skorB)
		if skorA < 0 || skorB < 0 {
			break
		}
		if skorA > skorB {
			fmt.Println("Hasil", i, ":", klubA)
			pemenang[jumlah] = klubA
			jumlah++
		} else if skorB > skorA {
			fmt.Println("Hasil", i, ":", klubB)
			pemenang[jumlah] = klubB
			jumlah++
		} else {
			fmt.Println("Hasil", i, ": Draw")
		}

		i++
	}
	fmt.Println("Pertandingan selesai")
	for j := 0; j < jumlah; j++ {
		fmt.Println("Hasil", nomor[j], ":", pemenang[j])
	}
}
