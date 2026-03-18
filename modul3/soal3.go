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
