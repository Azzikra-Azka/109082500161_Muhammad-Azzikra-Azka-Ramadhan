package main

import "fmt"

func jarakKuadrat(a, b, c, d int) int {
	return (a-c)*(a-c) + (b-d)*(b-d)
}

func didalam(cx, cy, r, x, y int) bool {
	if jarakKuadrat(cx, cy, x, y) <= r*r {
		return true
	}
	return false
}

func main() {
	var x1, y1, r1 int
	var x2, y2, r2 int
	var x, y int

	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&x, &y)

	d1 := didalam(x1, y1, r1, x, y)
	d2 := didalam(x2, y2, r2, x, y)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}