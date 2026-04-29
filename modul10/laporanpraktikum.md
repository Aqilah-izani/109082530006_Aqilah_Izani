# <h1 align="center">Laporan Praktikum Modul 10 - Array </h1>
<p align="center">Aqilah Izani - 109082530006</p>

## Unguided 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

#### soal1.go

```go
package main

import "fmt"

const MAX = 1000

func main() {
	var arr [MAX]float64
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	min := arr[0]
	max := arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println(min, max)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul10/output/ss_soal1.png)
Program tersebut digunakan untuk mendata anak kelinci yang akan dijual kepasar. pertama membuat array yang berkapasitas 1000, jadi program tersebut akan menampung sampai 1000 berat kelinci. Kemudian masuk kedalam func main. didala func main pengguna disuruh menginputkan bilangan, kemudian masuk kedalam for loop dimana di perulangan i diiterasikan dengan 0, dan i kurang dari n. kemudian nilai array dimasukan 1 per 1. Setelah array terisi, program mulai mencari nilai terkecil dan terbesar. yang dimana elemen pertama array yaitu arr[0], itu dijadikan sebagai nilai awal untuk variabel min dan max. kemudian masuk kedalam perulangan lagi dimana di dalam perulangan dimana perulangan dimulai dari index ke-1 hingga ke n. dan didalam perulangan tersebut terdapat if else yang diaman a jika nilai lebih kecila dari min maka nilai tersebut akan menggantikan min dan jika nilai lebih besar dari max maka nilai akan menggantikan max dan perulangan tersebut akan terus berlangsung sampai elemen dalam aarray di periksa kemudian jika sudah di periksa semua output yang di tampilkan adalah nilai terkecil dan nilai terbesar dalam elemen array. dan program pun selesai.

## Unguided 

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

#### soal2.go

```go
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

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul10/output/ss_soal2.png)
Program tersebut digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. pertama program membuat dua array yaiu arr dan hasil keduanya menyipan kapasitas  maksimum 1000 . kemudian pprogram tersebut membaca x dan y yang dimana pengguna di suruh menginputkan bilangan x dan bilangan y. yang dimana x menyatakan jumlah ikan yang akan di proses dan y menyatakan kapasitas dalam 1 wadah. kemudian program masuk kedalam perulangan. kemudian program masuk for index < x, diamana variabel index digunakan untuk menandai posisi ikan yang sedang di proses ke dalam array dan di mulai dari 0, dan variabel jumlahWadah digunakan untuk menghitung berapa banyak wadah yang sudah terisi. perulangan tersebut akan terus berjalan selama masih ada ikan yang belum di masukkan ke dalam wadah. kemudian di dalam perulangan tersebut ada variabel total untuk menjumlahkan berat ikan  dalam satu wadah, dan variabel hitung untuk menghitung berapa ikan yang sudah dimasukan ke wadah tersebut. kemudian terdapat peulangan lagi yaitu for hitung < y && index < x. dimana fungsi perulangan ini untuk memastikan jumlah ikan dalam satu wadah tidak melebihi kapasitas y dan tidak melewati jumlah ikan yang tersedia. didalam perulangan ini terdapat variabel index ditambah ke variabel total, kemudian index di naikan untuk berpindah ke ikan berikutnyam dan hitung dinaikan untuk mencatat jumlah ikan dalam wadah tersebut, setelah satu wadah penuh atau ikan habis, nilai total disimpan ke array hasil pada posisi jumlahWarda, lalu jumlahWadah ditambah satu. kemudian setelah semua ikan tersebut dikelompokan, program akan mencetak isi array hasil. kemudian program masuk kedapam penghitungan ratta-rata berat per wadah. dimana variabel penjumlahan digunakan untuk menjumlahkan semua nilai dalam array ahasil. setelah itu, nilai tersebut dibagi dengan jumlah wadah untuk mendapatkan rata-rata, kemudian menampilkan nilai rata-rata.


## Unguided 

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
#### soal3.go

```go
package main

import "fmt"

const NMAX = 100

type arrBalita [NMAX]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var min, max float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Masukan berat balita ke-", i+1, ": ")
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &min, &max)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata(data, n))
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul10/output/ss_soal3.png)
Program tersebut digunakan untuk t untuk mencatat berat balita, lalu menghitung berat minimum, maksimum, dan rata-rata. pertama program membuat array arrBalita dimana array tersebut berkapasitas 100 kemudian membuat fungsi hitungMinMax didalam fungsi tersebut pertama nilai awal minimun dan maximum diambil dari elemen pertama array. kemudian membuat pointer bMin dan bMax untuk mengakses nilai kemudian terdapat perulangan dari indeks 1 sampai n-1 dimana didalam perulangan tersebut terdapat perbandingan, jika lebih kecil dari *bMin maka dijadikan minumum baru dan jika lebih besar dari &bMax dijadikan maksimum baru. kemudian membuat fungsi rerata dimana fungsi tersebut di gunakan untuk menghitung rata-rata. Fungsi ini bekerja dengan menjumlahkan semua elemen array menggunakan variabel total, dan kemudian membagikan dengan n. kemudian menampilkan menggunakan print %.2f yang dimana %.2f agar angka ditampilkan dengan dua angka di belakang koma. dan programpun selesai 