package main

import "fmt"

func isprime(n int) bool{
	var f, i int
	f = 0
	i = 1
	for i <= n {
		if n % i == 0{
			f++
		}
		i++
	}
	return f == 2
}

func main() {
	var n,j int
	j = 0
	fmt.Scanln(&n)
	for n >= 0{
		if isprime(n) {
			j++
		}
		fmt.Scanln(&n)
	}
	fmt.Println("Banyak prima:", j)
}
