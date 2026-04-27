# <h1 align="center">Laporan Praktikum Modul 10 - Tugas Laprak Modul 10 </h1>
<p align="center">[Azka] - [109082500161]</p>

## Unguided 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.
#### 1.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Azzikra-Azka/109082500161_Muhammad-Azzikra-Azka-Ramadhan/blob/main/Modul10/output/Modul10Nomor1.png)
Program tersebut dibuat untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.

## Unguided 

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
#### 2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Azzikra-Azka/109082500161_Muhammad-Azzikra-Azka-Ramadhan/blob/main/Modul10/output/Modul10Nomor2.png)
Program tersebut buat menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.

## Unguided 

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
#### 3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arr arrBalita, n int, bMin *float64, bMax *float64) {
	*bMin = arr[0]
	*bMax = arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < *bMin {
			*bMin = arr[i]
		}
		if arr[i] > *bMax {
			*bMax = arr[i]
		}
	}
}

func rerata(arr arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int

	fmt.Print("Banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	var min, max float64
	hitungMinMax(data, n, &min, &max)
	rata := rerata(data, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Azzikra-Azka/109082500161_Muhammad-Azzikra-Azka-Ramadhan/blob/main/Modul10/output/Modul10Nomor3.png)
Program tersebut untuk mencatat data berat badan balita (dalam kg). Data tersebut akan dimasukkan ke dalam array.