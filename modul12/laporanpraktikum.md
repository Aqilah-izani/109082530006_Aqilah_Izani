# <h1 align="center">Laporan Praktikum Modul 12 - Selection Sort </h1>
<p align="center">Aqilah Izani - 109082530006</p>

## Unguided 

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing- masing daerah. Keterangan: Terdapat 3 daerah dalam contoh input, dan di masing-masing daerah mempunyai 5, 6, dan 3 kerabat.

#### soal1.go

```go
package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		selectionSort(rumah)
		for j := 0; j < m; j++ {
			fmt.Print(rumah[j])

			if j != m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul12/output/ss_soal1.png)
Program tersebut digunakan untuk mengurutkan nomor dari kecil hingga terbesar menggunakan algoritma selection sort. pertama membutat fungsi selectionsort dimana fungsi tersebut digunakan untuk mengurutkan array secara ascending yaitu dari kecil kebesar didalam fungsi tersebut pertama menentukan panjang array terlebuh dagulu dengan code "n := len(arr)" yang artinya lena(arr) menghitung jumlah array yang disimpan pada variabel n. kemudian membuat perulangan iterasi pertama dimulai dari indeks 0. karena mau mencari nilai terkecil terlebih dahulu maka menggunakan minIDX := i. kemudian terdapat perulangan lagi dimana perulangan ini berfungsi untuk mencari elemen  yang lebih kecil dari minIDX. jika ditemukan nilai terkecil maa minIDX akan diperbarui menjadi indeks elemen terkecil tersebut. Kemudian keluar dari perulangan kedua dan di dalam perulangan pertama terdapat code 'arr[i], arr[minIdx] = arr[minIdx], arr[i]' ini berfungsi untuk menukar elemen. jadi setelah elemen terkecil ditukar maka akan di ganti dengan posisi i. kemudian setelah keluar dari fungsi selectionsort buat fungsi main fungsi ini penting karena ini adalah tempat program dijalankan. didalam fungsi ini terdapat inputan jadi pengguna di suruh menginputkan sebuah bilangan, bilang tersebut berapa daerah yang ini diinputkan misal menginputkan 3 maka ada 3 daerah. setelah itu terdapat perulangan, perulang ini tujuannya untuk memproses setiap daerah satu persatu. kemudian pengguna disuruh menginputkan sebuah bilangan lagi, bilang ini artinya ada berapa nomor rumah. kemudian terdapat code 'rumah := make([]int, m)' code tersebut artinya membuat array dengan variabel m. kemudian terdapat perulangan lagi dan didalam perulangan tersebut pengguna disuruh menginputkan sebuah bilangan lagi, bilangan tersebut artinya membaca semua nomor rumah ke array. kemudian keluar dari perulangan tersebut kemudian urutkan array dengan menggunakan code 'selectionSort(rumah)'. Kemudian masuk ke perulangan lagi dimana perulangan ini untuk menamilkan hasilnya. 'for j := 0; j < m; j++ {
	fmt.Print(rumah[j])

	if j != m-1 {
		fmt.Print(" ")
	}
}' code tersebut digunakan untuk menampilkan isi array yang sudah terurut. dan Program selesai.
## Unguided 

### 2. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil. Format Masukan masih persis sama seperti sebelumnya. Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.

#### soal2.go

```go
package main

import "fmt"

func selectionSortAsc(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func selectionSortDesc(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		maxIdx := i

		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}

		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			var x int
			fmt.Scan(&x)

			if x%2 == 1 {
				ganjil = append(ganjil, x)
			} else {
				genap = append(genap, x)
			}
		}

		selectionSortAsc(ganjil)
		selectionSortDesc(genap)

		first := true

		for _, v := range ganjil {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(v)
			first = false
		}

		for _, v := range genap {
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(v)
			first = false
		}

		fmt.Println()
	}
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul12/output/ss_soal2.png)
Program tersebut digunakan untuk mengurutkan nomor rumah berdasarkan aturan tertentuu. jika nomor rumah ganjil maka dilakukan secara ascending jika nomor rumah genap makan akan mengurutkan secara descending. Pertama terdapat fungsi selectionsort asecending  dimana didalam fungsi ini terdapat perulangan dan perulangan tersebut digunakan untuk menentukan posisi elemen yang akan diisi nilai terkecil karena mau mencari nilai terkecil maka menggunakan minIDX kemudian program mencari nilai yang lebih kecil pada bagian array setelah indeks i menggunakan perulangan lagi didalam perulangan tersebut terdapat sebuah kondisi jika ditemukan nilai yang lebih kecil maka minIDX akan di perbarui menjadi ineks nilai terkecil tersebut setelah pencarian selesai dilakukan pertukaran posisi sehingga nilai terkecil beerpindah kedepan. dan proses tersebut akan terus berulang sampai array terurut dari kecil ke besar. Kemudian masuk kedalaman fungsi yang ke dua yaitu fungsi selection sort descending. ini digunakan untuk mengurutkan data dari besar ke kecil cara kerjanya hanpir sama dengan ascending tadi. Kemudian masuk kedalaam fungsi main dimana didalam fungsi main tersebut pengguna disuruh menginputkan sebuaah bilangan. bilangan tersebut artinya jumlah daerah jika menginputkan 3 maka ada 3 daerah yang akan di proses. Kemudian terdapat perulangan disini perogram akan melakukan perulangan sejumlah banyak daerah. di setipa daerah program membaca jumlah rumah, jumlah rumas tersebut di dapat dari inputan penggunaka kemudian membuat slice kosong. dimana slice ganjil digunakan untuk menyimpan nomor rumah ganjil sedangkan genap untuk nomor rumah genap. Selanjutnya program membaca semua nomor rumah dengan menggunakan perulangan. kemudian setiap anka di cek jika dimodulus 2= 1 maka ganjil dan masuk ke dalam array ganjil. jika tidak maka akan masuk kedalam array genap. jika array ganjil maka akan mengurutkan dari kecil kebesar jika array genap maka akan mengurutkan dari besar ke kecil.


## Unguided 

### 3. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah." Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313. Keluaran adalah median yang diminta, satu data per baris.

```go
package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func main() {
	var data []int
	var x int

	for {
		fmt.Scan(&x)
		if x == -5313 {
			break
		}
		if x == 0 {
			temp := make([]int, len(data))
			copy(temp, data)
			insertionSort(temp)

			n := len(temp)

			if n%2 == 1 {
				fmt.Println(temp[n/2])
			} else {
				median := (temp[n/2-1] + temp[n/2]) / 2
				fmt.Println(median)
			}

		} else {
			data = append(data, x)
		}
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul12/output/ss_soal3.png)
Program tersebut digunakan untuk mencarii median dari sekumpulan data. Pertama membuat fungsi insertion sort terlebih dahulu ini digunakan untuk mengurutkan array secara ascending yitu dari kecil kebesar. kemudian membuat perulangan di dalam fungsi ini, perulangan tersebut dimulai dari indeks 1 karena elemen pertama dianggap sudah terurut. kemudian terdapat code'key := arr[i]' variabel key ini adalah nilai yang akan si sisipkan ke posisi yang benar pada bagian array yang sudah terurut. kemudian periksa elemen sebelumnya menggunakan 'j:= i -1'. kemudian terdapat perulangan lagi dimana perulangan tersebut digunakan untuk menggeser elemen yang lebih besar dari key ke kanan. misal 7 23 11 saat 11 diproses maka 23 akan geser ke kanan sehingga menjadi 7 23 23 kemudian key dimasukan kembali dengan kode 'arr[j+1] = key' sehingga menjadi 7 11 23. Kemudian keluar dari fungsi insertion sort terdapat fungsi main di dalam fungsi main tedapat perulangan dan perulangan itu akan terus berjalan sampai di temukan angka - 5313. kemudian pengguna menginputkan sebuah bilangan jika yang di inputkan -51313 maka akan berhenti. pemberhentian tersebut menggunakan code 'break'. Jika yang di inputkan 0 maka program harus mencari median dari seluruh data yang sudah tersimpan sebelumnya. Pertama program membuat salnian array dengan 'temp := make([]int, len(data))
copy(temp, data)' code tersebut dilakukan agar data yang asli tidak berubah saat diurutkan. Kemudian data diurutkan setelah itu jumlah data disimpan. di dalam fungsi main tersebut juga terdapat sebuah kondisi dimana jika jumlah data ganjil maka median adalah elemen tengah dan jika jumlah data genap maka median dihitung dari rata-rata dua elemen tengah. Kemudian median di tampilakan dan jika input bukan 0 dan bukan - 5313 maka data akan di masukan ke array.