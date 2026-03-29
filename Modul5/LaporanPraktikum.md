# <h1 align="center">Laporan Praktikum Modul 5 - Tugas Laprak Modul 5 </h1>
<p align="center">[Azka] - [109082500161]</p>

## Unguided 

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.
#### Nomor1.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Print(" ")
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Azzikra-Azka/109082500161_Muhammad-Azzikra-Azka-Ramadhan/blob/main/Modul5/output/Modul5Nomor1.png)
Program tersebut dibuat untuk yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.

## Unguided 

### 2. yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.
#### Nomor2.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Print(" ")
	fmt.Scan(&n)

	bintang(n, 1)
}

func bintang(n int, i int) {
	if i > n {
		return
	}

	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	bintang(n, i+1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Azzikra-Azka/109082500161_Muhammad-Azzikra-Azka-Ramadhan/blob/main/Modul5/output/Modul5Nomor2.png)
Program tersebut yang digunakan untuk menampilkan pola bintang berikut ini dengan
menggunakan fungsi rekursif. N adalah masukan dari user.

## Unguided 

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang dibagi apa saja yang habis membagi N.
#### Nomor3.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Print(" ")
	fmt.Scan(&n)

	faktor(n, 1)
}

func faktor(n int, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Print(i, " ")
	}

	faktor(n, i+1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/Azzikra-Azka/109082500161_Muhammad-Azzikra-Azka-Ramadhan/blob/main/Modul5/output/Modul5Nomor3.png)
Program tersebut untuk mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N.