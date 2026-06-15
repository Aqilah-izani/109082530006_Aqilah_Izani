package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		datasisip := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > datasisip {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = datasisip
	}
}

func main() {
	var data []int
	var x int

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		data = append(data, x)
	}

	insertionSort(data)

	for i := 0; i < len(data); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(data[i])
	}
	fmt.Println()

	if len(data) <= 1 {
		fmt.Println("Data berjarak 0")
		return
	}

	jarak := data[1] - data[0]
	tetap := true

	for i := 2; i < len(data); i++ {
		if data[i]-data[i-1] != jarak {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
