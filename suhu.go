package main

import "fmt"

func Conversion(suhu float64) float64 {
	suhu = (float64(9) / float64(5) * suhu) + float64(32.0)
	return suhu
}

func main() {
	var suhu float64

	fmt.Print("Suhu dalam Celcius: ")
	fmt.Scanln(&suhu)
	fmt.Println("Suhu dalam Fahrenheit:", Conversion(suhu))
}
