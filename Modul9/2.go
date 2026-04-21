package main

import "fmt"

func main() {
	var a [100]int
	var n, x, idx, cari int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()

	for i := 1; i < n; i += 2 {
		fmt.Print(a[i], " ")
	}
	fmt.Println()

	for i := 0; i < n; i += 2 {
		fmt.Print(a[i], " ")
	}
	fmt.Println()

	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(a[i], " ")
		}
	}
	fmt.Println()

	fmt.Scan(&idx)
	for i := idx; i < n-1; i++ {
		a[i] = a[i+1]
	}
	n--

	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()

	sum := 0
	for i := 0; i < n; i++ {
		sum += a[i]
	}
	rata := float64(sum) / float64(n)
	fmt.Println(rata)

	var total float64
	for i := 0; i < n; i++ {
		s := float64(a[i]) - rata
		total += s * s
	}
	fmt.Println(total / float64(n))

	fmt.Scan(&cari)
	f := 0
	for i := 0; i < n; i++ {
		if a[i] == cari {
			f++
		}
	}
	fmt.Println(f)
}