# <h1 align="center">Laporan Praktikum Modul 2 - Tugas Laprak Modul 2 </h1>
<p align="center">[Azka] - [109082500161]</p>

## Unguided 

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### soal2A.go

```go
package main

import "fmt"

func main() {
	var(
		satu, dua, tiga string
		temp string
)
fmt.Print("Masukkan input string: ")
fmt.Scanln(&satu)
fmt.Print("Masukkan input string: ")
fmt.Scanln(&dua)
fmt.Print("Masukkan input string: ")
fmt.Scanln(&tiga)
fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

temp = satu
satu = dua
dua = tiga
tiga = temp
fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Azzikra-Azka/109082500161_Muhammad-Azzikra-Azka-Ramadhan/blob/main/LaprakModul2/output/Soal2A.png)
Program tersebut bertujuan untuk memindahkan nilai dari tiga variabel string secara bergantian

## Unguided 

### 2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.
#### soal2B.go

```go
package main

import "fmt"

func main() {

	var warna1, warna2, warna3, warna4 string
	berhasil := true

	for i := 1; i <= 5; i++ {

		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)

		if !(warna1 == "merah" && warna2 == "kuning" && warna3 == "hijau" && warna4 == "ungu") {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Azzikra-Azka/109082500161_Muhammad-Azzikra-Azka-Ramadhan/blob/main/LaprakModul2/output/Soal2B.png)
Program tersebut buat menerima input berupa warna dari ke 4 gelas reaksi
sebanyak 5 kali percobaan.

## Unguided 

### 3. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
#### soal2C.go

```go
package main

import "fmt"

func main() {

	var beratGram int
	var kg, sisa int
	var biayaKg, biayaSisa, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratGram)

	kg = beratGram / 1000
	sisa = beratGram % 1000

	biayaKg = kg * 10000

	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisa >= 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}

	total = biayaKg + biayaSisa

	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp.", total)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Azzikra-Azka/109082500161_Muhammad-Azzikra-Azka-Ramadhan/blob/main/LaprakModul2/output/Soal2C.png)
Program tersebut untuk menghitung biaya pengiriman berdasrkan berat parsel.