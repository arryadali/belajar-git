package main

import "fmt"

func MenghitungDisc(b float64, m bool) float64 {
	var harga float64
	if m == true && b > 100000 {
		harga = b - b*10/100

	} else if m == false && b > 100000 {
		harga = b - b*5/100
	} else {
		harga = b
	}
	return harga
}

func main() {
	var harga float64
	var membership bool
	fmt.Scan(&harga, &membership)
	fmt.Print("Total Bayar : ", MenghitungDisc(harga, membership))
}
