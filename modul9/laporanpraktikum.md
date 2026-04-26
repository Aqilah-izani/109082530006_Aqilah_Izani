# <h1 align="center">Laporan Praktikum Modul 3 - Fungsi </h1>
<p align="center">Aqilah Izani - 109082530006</p>

## Unguided 

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	tengah titik
	r      int
}

func jarak(p, q titik) float64 {
	dx := float64(p.x - q.x)
	dy := float64(p.y - q.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func didalam(c lingkaran, p titik) bool {
	return jarak(p, c.tengah) <= float64(c.r)
}

func main() {
	var l1, l2 lingkaran
	var p titik
	var in1, in2 bool

	fmt.Scan(&l1.tengah.x, &l1.tengah.y, &l1.r)
	fmt.Scan(&l2.tengah.x, &l2.tengah.y, &l2.r)
	fmt.Scan(&p.x, &p.y)

	in1 = didalam(l1, p)
	in2 = didalam(l2, p)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul9/output/ss_soal1.png)
Program tersebut digunakan untuk ngecek titik (x,y) apakah di dalam lingkaran 1, lingkaran 2, lingkaran 1 dan 2 atau di luar semua. pertama membuat struc dulu di beri type titik berfungsi untuk menyimpan koordinat x dan y. kemudianmembuat  tipe bentukan struck yaitu lingkaran yang berfungsi untuk menimpan titik tengah sebagai pusat lingkaran dan r sebagai radius. kemudian lanjut membuat fungsi jarak. Fungsi ini ada 2 titik yaitu p dan q. kemudian menghitung selisih dx dan dy terlebih dahulu utuk menghitung selisih dx yaitu px - qx dan untuk menghitung dy = py - qy. kemudian dx dan dy di kuadratkan dan di jumlahkan sehingga di peroleh jarak antar dua tituk. kemudian memnuat fungsi didalam, fungsi ini berfungsi untuk menentukan apakah sebuah titik berada dalam lingkaran. kemudian masuk ke func main pertama pengguna di suruh menginputkan koordinat pusat dan radius lingkaran, setelah itu pengguna disuruh menginputkan lagi koordinat ke dua dan radius lingkaran ke 2, kemudian pengguna disuruh menginputkan lagi koordinat titik. Jika kedua nilai in1 dan in2 bernilai true berarti titik berada di dalam ke dua lingkaran. jika hanya in1 yang true berarti hanya di dalam lingkaran pertama. jika in2 yang true berarti hanya di dalam lingkaran ke dua. jika keduanya fals berarti titik di luar kedua lingkaran. 

## Unguided 

### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:
a. Menampilkan keseluruhan isi dari array.
b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).
d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.
e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array
tersebut.

#### soal2.go

```go
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




```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul9/output/ss_soal2.png)
Program tersebut digunakan untuk menampung sekumpulan bilangan bulat. pertama buat func main di dalam func main terdapat beberapa code pertama tama mendeklarasikan sebuah array yaitu cost MAX = 100 var arr [MAX] int ini artinya array memiliki kapasitias tetap 100 elemen tidak berubah. kemudian pengguna menginptkan n, kemudian program membaca nilai n lalu mengisi array sebanya n elemen menggunakan perulangan. misal menginputkan 5 maka perulangan akan berjalan sampai ke 5. kemudian program menampilkan isi array. pertama semua elemen di tampilkan dari index 0 sampai n-1. kemudian program masuk ke menampilkan index ganjil saja, untuk menampilkan index ganjil menggunakan perulangan jika index dimodulus 2 == 1 berarti indexnya ganjil maka akan di tampilkan kemudian masuk ke code menampilkan index genap. untuk menampilkan index genap saja menggunakan perulangan jika index di modulus 2 = 0 maka termasuk genap. selanjutnya masuk ke program menampilkan elemen yang indexnya kelipatan x pertama pengguna memasukan inputan x terlebih dahulu disitu menggunakan perulangan. kemudian masuk ke program menghapus elemen index pertama pengguna menginputkan index berapa yang akan di hapus jika menginputkan 1 maka index satu akan di hapus. kemudian masuk program mencari rata- rata pertama semua elemen di jumlahkan terlebih dahulu ke dalam variaabel totaal, kemudian di bagi dengan n yang diinputkan tadi. karena hasil pembagian adalah desimal maka dikonversikan ke float64. kemudian masuk kedalam standar devisiasi . program tersebut menghitung selisih setiap elemen terhadap rata-rata lalu selisish tersebut dikuadratkan dan di jumlahkan. kemudian hasil total di bagi dengan n kemudian di akar kan menggunakan math.sqrt. kemudian masuk ke dalam program menghitung frequensi suatu bilangan. pertama program membaca sebuah nilai cari, kemudian melakukan perulangan untuk menghitung berapa nilai tersebut muncul di dalam array. setiap kali ditemukan elemen yang sama dengan cari, variabel freq akan di tambah satu.


## Unguided 

### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.

#### soal3.go

```go
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


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul9/output/ss_soal3.png)
Program tersebut digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. pertama terdapat array pemenang di gunakan untuk menyimpan nama klub pemenang. yang di mana ukuranya 100 dan 100 tersebut tidak bisa berubah. kemudian pengguna disuruh menginputkan klub A dan klub B. Setelah itu masuk ke perulangan for didalam perulangan for menampil kan pertandingan berapa kali yaitu dengan memasukan variabel i dimana variabel i ini dimulai dari 1. kemudian di dalam perulangan terdapat if else. jika salah satu skor kurang dari 0 atau berbentuk (-) maka perulangam akan berhenti, jika skor A lebih dari skor B maka klub A akan menang kemudian di cetak dan disimpan ke array pemenang pada index jumlah, dan jumlah itu akan di tambah satu, dan jika skor B lebih dari skor A maka klub B akan menang. dan jika skor nya sama maka hasilnya akan Draw. setelah keluar dari perulangan akan mencetak "pertandingan selesai" dan asuk ke bagian akhir yaitu for j := 0; j < jumlah; j++ {
		fmt.Println("Hasil", nomor[j], ":", pemenang[j])
	} ini berfungsi program menampilkan kembali hasil yangvtadi di hitung.

## Unguided 

### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.

#### soal4.go

```go
package main

import "fmt"

const NMAX = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var s string
	*n = 0

	for {
		fmt.Scan(&s)

		if s == "." || *n >= NMAX {
			break
		}

		t[*n] = rune(s[0])
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks : ")
	isiArray(&tab, &m)

	balikanArray(&tab, m)

	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul9/output/ss_soal4.png)
Program tersebut digunakan untuk untuk menampung sekumpulan karakter. pertama membuat array statis dengan kapasitas max 127 dengan tipe datanya rune, yaitu tipe untuk menyimpan karakter. kemudian membuat code fungsi isi array dimana di dalam fungsi tersebut mencantumkan variabel s bertipe data string dan menjumlahkan elemen awal dengan nol. kemudian terdapat perulangan juga di dalam perulangan, perulangan tersebut akan berhenti jika ketemu tanda titik. kemudian masuk ke dalam fungsi cetak array dimana di dalam fungsi tersebut setiap elemen dii cetak sebagai huruf. kemudian masuk kedalam fungsi balikanArray dimana fungsi ini adalah inti dari reverse di dalam fungsi tersebut terdapar perulangan yang dimana tujuannya elemen depan ditukar dengan elemen belakang kemudian bergerak ke tengah. kemudian membuat fungsi main di dalam fungsi main terdapat variabel tab dan n dimana tab tersebut array untuk menimpan huruf dan m untuk jumlah huruf yang akan di inputkan. kemudian program tersebut diminta untuk meminta pengguna menginputkan dan mengisi array. dan kemudian array tersebut di balik dan yang terakhir menampilkan hasil reverse . programpun selesai sampai disitu.

# Unguided 

### 5. Modifikasi program tersebut dengan menambahkan fungsi palindrom. Tambahkan instruksi untuk memanggil fungsi tersebut dan menampilkan hasilnya pada program utama.

#### soal4modif.go

```go
package main

import "fmt"

const NMAX = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var s string
	*n = 0

	for {
		fmt.Scan(&s)

		if s == "." || *n >= NMAX {
			break
		}

		t[*n] = rune(s[0])
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks : ")
	isiArray(&tab, &m)

	fmt.Println("Palindrom ?", palindrom(tab, m))

	balikanArray(&tab, m)

}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul9/output/ss_modifsoal4.png)
Program tersebut adalah modifikasi dari soal nomer 4 dimana hanya menambahkanapakah palindom atau bukan. pertama membuat array statis dengan kapasitas max 127 dengan tipe datanya rune, yaitu tipe untuk menyimpan karakter. kemudian membuat code fungsi isi array dimana di dalam fungsi tersebut mencantumkan variabel s bertipe data string dan menjumlahkan elemen awal dengan nol. kemudian terdapat perulangan juga di dalam perulangan, perulangan tersebut akan berhenti jika ketemu tanda titik. kemudian masuk ke dalam fungsi cetak array dimana di dalam fungsi tersebut setiap elemen dii cetak sebagai huruf. kemudian masuk kedalam fungsi balikanArray dimana fungsi ini adalah inti dari reverse di dalam fungsi tersebut terdapar perulangan yang dimana tujuannya elemen depan ditukar dengan elemen belakang kemudian bergerak ke tengah.kemdia masuk kedalam fungsi palindrom didalam fungsi tersebut terdapat perulangan dimana di dalam perulangan membandingkan elemen dari depan dengan elemen belakang dan indeks i di bandingkan dengan n-1-i kalau ada satu yang tidak sama hasilnya false kalau semua sama maka true. kemudian membuat fungsi main di dalam fungsi main terdapat variabel tab dan n dimana tab tersebut array untuk menimpan huruf dan m untuk jumlah huruf yang akan di inputkan. kemudian memanggil fungsi paindrom dan menampilkan hasil dari palindorm tersebut. programpun selesai sampai disitu.

