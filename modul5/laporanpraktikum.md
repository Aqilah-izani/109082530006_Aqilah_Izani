# <h1 align="center">Laporan Praktikum Modul 5 - Rekursif </h1>
<p align="center">Aqilah Izani - 109082530006</p>

## Unguided 

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.

#### soal1.go

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	var n, i int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	for i = 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul5/output/ss_soal1.png)
Program tersebut digunakan untuk menghitung bilangan fibionacci dari 0 sampai n, n tersebut nanti yang akan kita inputkan nanti misal kita inputkan n =  10 maka dari 0 - ke 10 tersebut. pertama buat fungsi fibonacci terlebih dahulu kemudian didalam fung tersebut terdapat sebuah kondisi memakai if else, dimana jika n = 0 maka akan menghasilkan 0, jika n = 1 maka akan menghasilkan 1 tetapi jika 0 nya bukan 0 dan 1 maka akan masuk ke code return fibonacci(n-1) + fibonacci(n-2) ini artinya mengambil nilai sebelumnya kemudian di jumlahkan misal fibionanci(4) maka fibionancci 3 + fibionanci(2) kemudian lanjut ke fibionaci(3) maka fibionaci(2)+fibionaci(1) begitu terus sampai menemukan fibionaci(1) dan fibionaci(0). kemudian masuk kedalan func main didalam func main pengguna disuruh untuk menginputkan bilangan n dan didalam func main juga terdapat perulangan.kemudian di perulangan program tersebut akan mengulang dari 0 sampai ke n misal menginputkan n nya 10 maka program akan mengulang sebanyak 10 kali setiap mengeprint nilai i maka akan memanggil funsi fibionaci tadi. contoh kalau program di jalankan pengguna disuruh menginputkan n = 10 maka bilangan fibionaci nya adalah 0 1 1 2 3 5 8 13 21 34 55. hitung dari 0 dulu maka dia akan menampilkan 0 nahh kemudian nambah jadi 1 maka bilangan fibionacinya 1 kemudian i tambah 1 lagi jadi 2 berarti ini masuk ke rumus jadi 0 + 1= 2 dann seterusnya .

## Unguided 

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.

#### soal2.go

```go
package main

import "fmt"

func bintang(n int) {
	if n == 1 {
		fmt.Println("*")
	} else {
		bintang(n - 1)

		for i := 0; i < n; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	bintang(n)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul5/output/ss_soal2.png)
Program tersebut digunakan untuk menghitung pola pada bintang pertama bikin terlebih dahulu fungsi bintang didalam fungsi bintang terdapat terdapat kondisi if else jika n = 1 maka akan menampilkan 1 bintang (*), jika n lebih dari 1 maka akan masuk ke code bintang(n - 1) kemudian didalam kondisi if else juga ada perulangan yang dimana perulangan tersebut dimulai dari 0 sampai n. disitu akan menampilkan bintang sebanyak n kali misal n = 3 maka bintang(3) akan menampilkan 2 bintang kemudian 1 bintang lalu di balik ke atas lalu di kembalikan lagi katas jadi terkecil menampilkan bintang 1, bintang 2 dan bintang 3. Kemudian masuk ke fungsi main di dalam fungsi main pengguna disuruh memasukan 1 input angka kemudian menampilkan n

## Unguided 

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
#### soal3.go

```go
package main

import "fmt"

func faktor(n int, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Print(i, " ")
	}

	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)
	faktor(n, 1)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul5/output/ss_soal3.png)
Program tersebut digunakan untuk menampilkan faktor untuk bilangan suatu N. pertama bikin func faktor terlebih dahulu didalam bilangan faktor terdapat suatu kondisi if else, dimana jika i lebih dari n maka akan berhenti di cetak, dan jika n sisa bagi = 0 maka i akan langsung di tampilkan. setelah itu di luar kondisi if else, fungsi akan memanggil dirinya. kemudian masuk ke dalam func main pengguna disuruh menginputkan n, kemudian code faktor(n, 1) itu artinya i dimulai dari angka 1.