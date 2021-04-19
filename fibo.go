package main

import "fmt"

func fibo(n int) int {
	var n1, n2, i, x, y int
	i = 0

	for i < n {
		if i == 0 {
			x = 0
			n1 = n2
			n2 = x
		} else if i == 1 {
			x = 1
			n1 = n2
			n2 = x
		} else {
			x = n1 + n2
			n1 = n2
			n2 = x
		}
		y += x
		i++
	}
	return y
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(fibo(n))
}
