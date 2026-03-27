package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0

	var waktu int
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama string

	fmt.Scan(&nama)

	// peserta pertama jadi acuan
	var maxSoal, minWaktu int
	var pemenang string

	if nama != "Selesai" {
		hitungSkor(&maxSoal, &minWaktu)
		pemenang = nama
	}

	fmt.Scan(&nama)

	for nama != "Selesai" {
		var soal, skor int
		hitungSkor(&soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			maxSoal = soal
			minWaktu = skor
			pemenang = nama
		}

		fmt.Scan(&nama)
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}