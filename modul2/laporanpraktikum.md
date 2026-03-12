# <h1 align="center">Laporan Praktikum Modul 2 - review alpro 1 </h1>
<p align="center">Aqilah Izani - 109082530006</p>

## Unguided 

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul2/output/ss_soal1.png)
program tersebut digunakan untuk menukar urutan string. pertama masukan tiga variabel string yaitu variabel satu, dua, dan tiga kemudian ada variabel temp nah variabel temp tersebut digunakan untuk tempat sementara. setelah itu, program akan menampilkan "masukan input string: " setelah itu pengguna disuruh untuk menginputkan variabel satu, program menampilkan "masukan input string: " lagi dan pengguna menginputkan variabel dua, dan program menampilkan "masukan input string: " lagi dan menginputkan variabel tiga. setelah itu akan menampilkan output awal setelah itu penukaran dimulai program menukar nilai string dengan cara menimpan nilai satu disimpan pada temp kemudian mengubah nilai satu menjadi dua, nilai dua menjadi tiga dan tiga menjadi nilai yang disimpan pada temp. kemudian program menampilkan output akhir dan program selesai.


### 2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.
#### soal2.go

```go
package main

import "fmt"

func main() {
	var p1, p2, p3, p4 string
	var berhasil bool
	var i int

	berhasil = true
	for i = 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&p1, &p2, &p3, &p4)
		if p1 != "merah" || p2 != "kuning" || p3 != "hijau" || p4 != "ungu" {
			berhasil = false
		}
	}
	fmt.Print("BERHASIL: ", berhasil)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul2/output/ss_soal2.png)
Program tersebut digunakan untuk menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 
kali percobaan. Kemudian program akan menampilkan true ketika urutan warna sesuai dengan 
informasi yang diberikan, dan akan false ketika warnanya tidak sesuai. Pertama 
program tersebut berhasil = true kemudian program masuk ke dalam perulangan didalam 
peruulangan i diiterasi kan 1 kemudian i tersebut akan bertambah sampai <= 5, Karena kita akan 
memasukan 5 kali percobaan. Kemudian untuk menghasilkan ouput percobaan 1 maka 
menggunakan print i kemudian pengguna menginputkan 4 warna padaa percobaan tersebut jika 
inputan warna pertama tidak sama dengan merah, warna kedua tidak sama dengan kuning, warna 
ketiga tidak sama dengan hijau, dan warna keempat tidak sama dengan  ungu maka berhasil = false 
tetapi jika warna tersebut merah kuning hijau ungu maka berhasil = true, perulangan tersebut akan terus berjalan sampai ke percobaan ke 5 setelah perulangan selesai akan mengeprint “Berhasil: ”, berhasil setelah itu program selesai.


## Unguided 

### 3. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.

```go
package main

import "fmt"

func main() {
	var Berat_parsel, Kg, gr, biaya_jasa, total_biaya, detail_biaya int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&Berat_parsel)

	Kg = Berat_parsel / 1000
	gr = Berat_parsel % 1000
	biaya_jasa = Kg * 10000

	if gr < 1 {
		detail_biaya = gr * 0
	} else if gr >= 500 {
		detail_biaya = gr * 5
	} else {
		detail_biaya = gr * 15
	}
	total_biaya = biaya_jasa + detail_biaya
	fmt.Println("Detail berat : ", Kg, "Kg", "+", gr, "gr")
	fmt.Println("Detail biaya : ", "RP.", biaya_jasa, "+", "Rp.", detail_biaya)
	fmt.Println("Total biaya: Rp.", total_biaya)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Aqilah-izani/109082530006_Aqilah_Izani/blob/main/modul2/output/ss_soal3.png)
Program tersebut digunakan untuk untuk menghitung total biaya jasa pengiriman. Nahh pertama program tersebut disuruh untuk menginputkan berat parsel terlebih dahulu. Kemudian, berat Kg dan grdipisah  memakai pembagian dan modulus. Setelah itu Kg dan gr terpisah program tersebut akan menghitung jasa pengiriman Kg terlebih dahulu yang di mana perKg nya 10.000 program tersebut ditulis Kg * 10.000. setelah itu menghitung biaya tambahan jasa pengiriman yang berada di gr memakai if else. Jika kurang dari 1 Kg tidak ada penambahan jasa pengiriman. Jika >= 500 gr bakal dikali 5 perak dan jika <= 500 gr bakal dikali 15 perak
