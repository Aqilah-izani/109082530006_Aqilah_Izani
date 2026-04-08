package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0

	var waktu, i int
	for i = 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal = *soal + 1
			*skor += waktu
		}
	}
}
func main() {
	var nama, pemenang string
	var maxSoal, minSkor int
	fmt.Scan(&nama)
	for nama != "Selesai" {
		var soal, skor int
		hitungSkor(&soal, &skor)
		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		}
		fmt.Scan(&nama)
	}
	fmt.Println(pemenang, maxSoal, minSkor)
}
