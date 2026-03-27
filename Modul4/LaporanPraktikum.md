# <h1 align="center">Laporan Praktikum Modul 4 - Tugas Laprak Modul 4 </h1>
<p align="center">[Azka] - [109082500161]</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p). Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut!
#### Nomor1.go

```go
package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, hasil *int) {
	var fn, fnr int
	factorial(n, &fn)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}

func kombinasi(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	var p1, c1, p2, c2 int

	permutasi(a, c, &p1)
	kombinasi(a, c, &c1)

	permutasi(b, d, &p2)
	kombinasi(b, d, &c2)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Azzikra-Azka/109082500161_Muhammad-Azzikra-Azka-Ramadhan/blob/main/Modul4/output/Modul4Nomor1.png)
Program tersebut dibuat untuk mempelajari kombinasi dan permutasi dari mata kuliah matematika diskrit dan diimplementasikan ke suatu program.

## Unguided 

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.
#### Nomor2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Azzikra-Azka/109082500161_Muhammad-Azzikra-Azka-Ramadhan/blob/main/Modul4/output/Modul4Nomor2.png)
Program tersebut dibuat untuk mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.