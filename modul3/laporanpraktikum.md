# <h1 align="center">Laporan Praktikum Modul 3 - Fungsi </h1>
<p align="center">Aqilah Izani - 109082530006</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas?Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi,dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi aterhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.


#### soal1.go

```go
package main

import "fmt"

func faktorial(n int) int {
	var hasil, i int
	hasil = 1
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutation(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func combination(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul3/output/ss_soal1.png)
Program tersebut digunakan untuk menghitung permutasi dan kombinasi. Pertama bikin fungsi faktorial dulu didalam fungsi faktrorial terdapat perulangan yang dimana hasil dimulai dengan 1 kemudian melakukan perulangan dari 1 sampai ke n. setelah itu masuk ke fungsi permutasi didalam fungsi permutasi terdapat rumus untuk menghitung permutasi yaitu faktorial(n)/ faktorial(n-r). Kemudian masuk ke fungsi combination didalam fungsi combination terdapat rumus  faktorial(n) / (faktorial(r) * faktorial(n-r)). setelah itu masuk ke main kembudian pengguna menginputkan 4 angka yang di simpan dalam variabel a, b, c, d. setelah itu program mengeprint baris pertama adalah hasil permutasi dan kombinasi dari a terhadap c, kemudian baris ke dua mengeprint hasil permutasi dan kombinasi dari b terhadap d. dan program selesai.

## Unguided 

### 2. Diberikan tiga buah fungsi matematika yaitu f (x) = x^2, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g( (x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function. Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi. Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!


#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func fogoh(x int) int {
	return f(g(h(x)))
}

func gohof(x int) int {
	return g(h(f(x)))
}

func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	fmt.Println(fogoh(a))
	fmt.Println(gohof(b))
	fmt.Println(hofog(c))
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul3/output/ss_soal2.png)
Program tersebut digunakan untuk menghitung fogoh(a), gohof(b), hofog(c). pertama buat 3 fungsi dasar terlebih dahulu yaitu, f(x), g(x), h(x) yang di mana fungsi f(x) memiliki rumus x^2, g(x) = x -2 dan h(x) = x +1 . kemudian membuat 3 fungsi lagi untuk menghitung fogoh,gohof,hofog. yang di mana rumus dari fogoh = f(g(h(x))), gohof = g(h(f(x))), hofog = h(f(g(x))). kemudian masuk ke func main di dalam func main pengguna menginputkan 3 angka yang di dalam variabael a,b,c. kemudian program mengeprint fogoh(a),gohof(b), hofog(c). cara menghitungnya seperti ini misal pengguna meng inputkan 7 2 10 yang artinya a = 7, b = 2, dan c = 10 kemudian program print fogoh (a), jadi rumus dari fogoh f(g(h(x))) hitung yang paling dalam dulu  h(x)= 7+1 = 8 kemudian gitung g(x)= 8-2= 6 dan f(x)= 6*6 = 36 program akan menapilkan 36 dan program pun selesai. 


## Unguided 

### 3. [Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".


#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}
func dalamLingkaran(cx, cy, r, x, y float64) bool {
	d := jarak(cx, cy, x, y)
	return d <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64
	var titik1, titik2 bool

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	titik1 = dalamLingkaran(cx1, cy1, r1, x, y)
	titik2 = dalamLingkaran(cx2, cy2, r2, x, y)

	if titik1 && titik2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if titik1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if titik2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul3/output/ss_soal3.png)
Program tersebut digunakan untuk ngecek titik (x,y) apakah di dalam lingkaran 1, lingkaran 2, lingkaran 1 dan 2 atau di luar semua. pertama membuat fungsi jarak yang dimna rumusnya akar(a-c)*(a-c)+(b-d)*(b-d). kemudian membuat fungsi didalam lingkaran. nahh didalam fungsi lingkaran program memanggil fungsi jarak setelah didapat jaraknya program akan membandingkan jarak sama radius. kalau jaraknya lebih kecil atau sama dengan radius berarti titiknya masih di dalam limglaran. kalau lebih besar, berarti sudah di luar. kemudian masuk ke fungsi main kemudian pengguna menginputkan cx1,cy1,r1 kemudian pengguna menginputkan lagi cx2,cy2, r2 kemuidan menginputkan x dan y. kemudian masuk ke else if. jika titik1 dan titik 2 true akan mencetak didalam lingkaran 1 dan 2 tetapi jika titik1 bernilai true maka titik di dalam lingkaran 1. jika titik2 bernilai true maka akan mencetak titik 2 tetapi jika ke duanya false akan mencetak titik di luar lingkaran 1 dan 2.