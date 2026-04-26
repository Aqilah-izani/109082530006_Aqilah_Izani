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
