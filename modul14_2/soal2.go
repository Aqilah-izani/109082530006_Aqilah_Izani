package main

import "fmt"

type Buku struct {
	ID        string
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

func DaftarkanBuku(pustaka *[]Buku, n int) {
	for i := 0; i < n; i++ {
		var b Buku

		fmt.Scan(
			&b.ID,
			&b.Judul,
			&b.Penulis,
			&b.Penerbit,
			&b.Eksemplar,
			&b.Tahun,
			&b.Rating,
		)

		*pustaka = append(*pustaka, b)
	}
}

func CetakTerfavorit(pustaka []Buku, n int) {
	maxIdx := 0

	for i := 1; i < n; i++ {
		if pustaka[i].Rating > pustaka[maxIdx].Rating {
			maxIdx = i
		}
	}

	fmt.Println(
		pustaka[maxIdx].Judul,
		pustaka[maxIdx].Penulis,
		pustaka[maxIdx].Penerbit,
		pustaka[maxIdx].Tahun,
	)
}

func UrutBuku(pustaka []Buku, n int) {
	for i := 1; i < n; i++ {

		temp := pustaka[i]
		j := i - 1

		for j >= 0 && pustaka[j].Rating < temp.Rating {
			pustaka[j+1] = pustaka[j]
			j--
		}

		pustaka[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka []Buku, n int) {

	batas := 5

	if n < 5 {
		batas = n
	}

	for i := 0; i < batas; i++ {
		fmt.Println(pustaka[i].Judul)
	}
}

func CariBuku(pustaka []Buku, n int, r int) {

	kiri := 0
	kanan := n - 1

	for kiri <= kanan {

		tengah := (kiri + kanan) / 2

		if pustaka[tengah].Rating == r {

			fmt.Println(
				pustaka[tengah].Judul,
				pustaka[tengah].Penulis,
				pustaka[tengah].Penerbit,
				pustaka[tengah].Tahun,
				pustaka[tengah].Eksemplar,
				pustaka[tengah].Rating,
			)

			return
		}

		if r > pustaka[tengah].Rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	fmt.Println("Tidak ada buku dengan rating seperti itu")
}

func main() {

	var n int
	fmt.Scan(&n)

	var pustaka []Buku

	DaftarkanBuku(&pustaka, n)

	CetakTerfavorit(pustaka, n)

	UrutBuku(pustaka, n)

	Cetak5Terbaru(pustaka, n)

	var ratingCari int
	fmt.Scan(&ratingCari)

	CariBuku(pustaka, n, ratingCari)
}
