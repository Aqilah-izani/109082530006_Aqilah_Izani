package main

import (
	"fmt"
	"math"
)

func main() {
	const MAX = 100
	var arr [MAX]int
	var n, total, idx, x int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// a. menampilkan semua isi array
	fmt.Println("Semua elemen:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// b. indeks ganjil saja
	fmt.Println("Indeks ganjil:")
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// c. indeks genap saja
	fmt.Println("Indeks genap:")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// d. kelipatan x
	fmt.Scan(&x)
	fmt.Println("Indeks kelipatan", x, ":")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// e. hapus elemen di indeks tertentu
	fmt.Scan(&idx)
	for i := idx; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	fmt.Println("Setelah dihapus:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// f. rata-rata
	total = 0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	rata := float64(total) / float64(n)
	fmt.Println("Rata-rata:", rata)

	// g. standar deviasi
	var jumlah float64
	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - rata
		jumlah += selisih * selisih
	}
	std := math.Sqrt(jumlah / float64(n))
	fmt.Println("Standar deviasi:", std)

	// h. frekuensi suatu bilangan
	var cari int
	fmt.Scan(&cari)

	freq := 0
	for i := 0; i < n; i++ {
		if arr[i] == cari {
			freq++
		}
	}
	fmt.Println("Frekuensi:", freq)
}
