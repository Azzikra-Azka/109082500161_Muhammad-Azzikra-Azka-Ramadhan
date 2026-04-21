package main

import "fmt"

func main() {
	var klubA, klubB string
	var hasil [100]string
	var jumlah int = 0
	var skorA, skorB int

	fmt.Scan(&klubA, &klubB)
	fmt.Scan(&skorA, &skorB)

	for skorA >= 0 && skorB >= 0 {
		if skorA > skorB {
			hasil[jumlah] = klubA
		} else if skorB > skorA {
			hasil[jumlah] = klubB
		} else {
			hasil[jumlah] = "Imbang"
		}
		jumlah++

		fmt.Scan(&skorA, &skorB)
	}

	for i := 0; i < jumlah; i++ {
		fmt.Println("Hasil", i+1, ":", hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}