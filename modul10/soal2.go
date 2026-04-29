package main

import "fmt"

const MAX = 1000

func main() {
	var arr [MAX]float64
	var hasil [MAX]float64
	var x, y int

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&arr[i])
	}

	index := 0
	jumlahWadah := 0

	for index < x {
		total := 0.0
		hitung := 0

		for hitung < y && index < x {
			total += arr[index]
			index++
			hitung++
		}

		hasil[jumlahWadah] = total
		jumlahWadah++
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Print(hasil[i], " ")
	}
	fmt.Println()

	penjumlahan := 0.0
	for i := 0; i < jumlahWadah; i++ {
		penjumlahan += hasil[i]
	}

	rata := penjumlahan / float64(jumlahWadah)
	fmt.Println(rata)
}
