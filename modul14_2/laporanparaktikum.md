# <h1 align="center">Laporan Praktikum Modul 12 - Selection Sort </h1>
<p align="center">Aqilah Izani - 109082530006</p>

## Unguided 

### 1. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array. Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".

#### soal1.go

```go
package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		datasisip := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > datasisip {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = datasisip
	}
}

func main() {
	var data []int
	var x int

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		data = append(data, x)
	}

	insertionSort(data)

	for i := 0; i < len(data); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(data[i])
	}
	fmt.Println()

	if len(data) <= 1 {
		fmt.Println("Data berjarak 0")
		return
	}

	jarak := data[1] - data[0]
	tetap := true

	for i := 2; i < len(data); i++ {
		if data[i]-data[i-1] != jarak {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul14_2/output/ss_soal1.png)
Program tersebut digunakan untuk membaca sekumpulan bilangan bulat yang di akhiri dengan bilangan negarif. dimana perogram tersebut memanfaatkan insertion sort. Jadi bilangan positif yang diinputkan akan dimasukan kedalam array dan bilangan negatif digunakan sebagai penanda akhir input dan tidak tersimpan ke dalam array. setelah seluruh data tersimpan, array diurutkan menggunakan algoritma Insertion Sort secara ascending dari kecil ke besar. Insertion sort bekerja dengan cara mengambil satu elemn kemudian menyisipkan ke posisi yang tepat pada bagian array yang sudah terurut. Proses tersebut akan terus berulang sampai elemen berada pada urutan yang benar. Setelah proses pengurutan selesai, program menampilkan isi array yang telah terurut. Selanjutnya program memeriksa apakah selisih antara setiap pasangan elemen yang berurutan memiliki nilai yang sama. Nilai selisih pertama digunakan sebagai acyan kemudian dibandingkan dengan selisih elemen berikutnya. Jika selih bernilai sama maka program menampilkan pesan "Data berjarak x" dengan x merupakan nilai selisih tersebut. Jika ditemukan selisih yang berbeda, maka program menampilkan pesan "Data berjarak tidak tetap".
## Unguided 

### 2. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:
const nMax : integer = 7919
type Buku = <
id, judul, penulis, penerbit : string
eksemplar, tahun, rating : integer >
type DaftarBuku = array [ 1..nMax] of Buku
Pustaka : DaftarBuku
nPustaka: integer
Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang
menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya,
masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris
terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.
Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua
adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku
yang dicari sesuai rating yang diberikan pada masukan baris terakhir.

#### soal2.go

```go
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


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul14_2/output/ss_soal2.png)
Program tersebut digunakan untuk mengelola data buku pada sebuah perpustakaan. pertama membuat struct dengan type buku dan di dalam struct tersebut terdapat ID, judul,penulis dan penerbit yang menggunakan tipe data string dan eksemplar, tahun dan rating menggunakan tipe data integer. Kemudian struct buku digunakan untuk menyimpan seluruh informasi yang berkaitan dengan satu buku. penggunaan struct bertujuan agar mudah di kelola. kemduidan terdapat fungsi CetakTerfavorit digunakan untuk mencari buku dengan reting tertinggi. Program melakukan pencarian seluruh elemen pada slice dan menyimpan indeks buku yang memiliki rating terbesar. Setelah proses pencarian selesai, informasi buku tersebut ditampilkan berupa judul, penulis, penerbit dan tahun terbit.
Kemudian terdapat fungsi Urut Buku dimana fungsi ini berfungsi untuk mengurutkan data buku berdasarkan rating secara descending artinya dari besar ke kescil. Kemudiang masuk ke Prosedur Cetak5Terbaru diamana prosedur tersebut menampilkan maksimal lima buku dengan rating  tertinggi. Karena data telah diurutkan secara menurun berdasarkan rating menggunakan binary serch kemudian masuk ke fungsi main diama fungsi ini merupakan program utama yang mengatur jalannya program dimana langkahnya membaca jumlah buku, memanggil prosedur DaftarBuku, Menampilkan buku terfavorit menggunakan CetakTerfavorit mengurutkan data menggunakan UrutBuku, Menampilkan lima buku dengan rating tertinggi menggunakan Cetak5Terbaru, Membaca rating yang ingin di cari dan melakukan pencarian .