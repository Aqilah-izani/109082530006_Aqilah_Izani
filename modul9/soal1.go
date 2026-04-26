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
