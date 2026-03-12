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
