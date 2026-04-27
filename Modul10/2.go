package main

import "fmt"

type arr [1000]float64

func main() {
	var ikan arr
	var totalWadah arr
	var x, y int

	fmt.Print("Jumlah ikan: ")
	fmt.Scan(&x)

	fmt.Print("Kapasitas per wadah: ")
	fmt.Scan(&y)

	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	idx := 0
	jumlahWadah := 0

	for idx < x {
		totalSementara := 0.0
		jumlahIsiWadah := 0

		for jumlahIsiWadah < y && idx < x {
			totalSementara += ikan[idx]
			idx++
			jumlahIsiWadah++
		}

		totalWadah[jumlahWadah] = totalSementara
		jumlahWadah++
	}

	fmt.Println("Total berat tiap wadah:")
	totalSemua := 0.0

	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%.2f ", totalWadah[i])
		totalSemua += totalWadah[i]
	}

	rata := totalSemua / float64(jumlahWadah)
	fmt.Printf("\nRata-rata per wadah: %.2f\n", rata)
}