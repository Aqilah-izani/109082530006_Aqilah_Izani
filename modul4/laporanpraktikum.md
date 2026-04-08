# <h1 align="center">Laporan Praktikum Modul 4 - Prosedur </h1>
<p align="center">Aqilah Izani - 109082530006</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p) Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.

#### soal1.go

```go
package main

import "fmt"

func factorial(n int, hasil *int) {
	var i int
	*hasil = 1
	for i = 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var faktorialn, faktorialnr int
	factorial(n, &faktorialn)
	factorial(n-r, &faktorialnr)
	*hasil = faktorialn / faktorialnr
}

func combination(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	var p1, c1, p2, c2 int
	fmt.Print("Masukan bilangan a, b, c, dan d: ")
	fmt.Scan(&a, &b, &c, &d)
	permutation(a, c, &p1)
	combination(a, c, &c1)
	permutation(b, d, &p2)
	combination(b, d, &c2)
	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul4/output/ss_soal1.png)
Program tersebut digunakan untuk menghitung permutasi dan kombinasi. Pertama bikin fungsi faktorial dulu didalam fungsi faktrorial terdapat perulangan yang dimana hasil dimulai dengan 1 kemudian kemudian terdapat perulangan dari 1 sampai ke n,nahh kemudian disitu code tersebut akan melakukan perkalian dari 1 sampai ke angka n tersebut misal masukin 3 maka dia akan menghitung 1x2x3 sampai dapat 6. setelah itu masuk ke fungsi permutasi didalam fungsi permutasi terdapat code factorial(n, &faktorialn) itu artinya faktorial = n! dan ada juga code factorial(n-r, &faktorialnr) itu artinya faktorial = n-r!.kemudian karena rumus dari permutasi = n!/(n-r)! maka codenya jadi *hasil = faktorialn / faktorialnr. Kemudian masuk ke fungsi kombinasi didalam fungsi terdapat code factorial(n, &fn) itu artinya faktorial = n!, factorial(r, &fr) artinya = faktorial r dan code factorial(n-r, &fnr) artinya faktorial = (n-r)!, karena rumus cominasi = n!/r!(n-r)! maka code yang di tulis fn / (fr * fnr). setelah itu masuk ke func main kemudian pengguna menginputkan 4 angka yang di simpan dalam variabel a, b, c, d. kemudian program mengeprin permutasi dari a dan c setelah itu mengeprint kombinasi a dan c setelah itu mengeprint permutasi b dan d kemudian print lagi kombinasi b dan d. program pun selesai.

## Unguided 

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. prosedure hitungSkor(in/out soal, skor : integer) Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.


#### soal2.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul4/output/ss_soal2.png)
Program tersebut digunakan untuk menghitung skor nilai suatu lomba. pertama bikin fungsi hitungskor terlebih dahulu yang di mana di dalam fungsi tersebut ada code *soal = 0 dan *skor = 0 itu artinya belum ada soal yang berhasil yaitu masih 0 dan belum ada waktu yang dihitung masih 0, kemudian di dalam fungsi hitungskor juga terdapat variabel waktu, dan i yang ber tipe data integer kemudian juga terdapat perulangan  i<8 itu karena setiap peserta mempunyai 8 soal. kemudian masuk ke waltu jika waktu <301 menit maka soal berhasil masuk ke *soal= *soal + 1 artinya jumlah soal akan ditambah 1 dan total waktu juga ditambah. Kemudian masuk ke func main di func main pengguna menginputkan nama terlebih dahulu kemudian masuk ke perulangan untuk for nama != "Selesai" artinya program tersebut akan terus berjalan sampai menemukan kata Selesai tetapi jika belum menemukan kata Selesai akan masuk ke hitungSkor(&soal, &skor) nahh kemudian jika soal lebih banyak akan menang atau kalau soal sama melihat waktu yang lebih cepat jiaka menang akan print nama, soal, dan skor. program selesai.