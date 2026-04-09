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
