package main

import "fmt"

type arr [1000]float64

func main() {
	var data arr
	var n int

	fmt.Print(" ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Berat ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	min := data[0]
	max := data[0]

	for i := 1; i < n; i++ {
		if data[i] < min {
			min = data[i]
		}
		if data[i] > max {
			max = data[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}