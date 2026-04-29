package main

import "fmt"

const MAX = 1000

func main() {
	var arr [MAX]float64
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	min := arr[0]
	max := arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println(min, max)
}
