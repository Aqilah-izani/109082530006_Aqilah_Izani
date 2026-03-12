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
