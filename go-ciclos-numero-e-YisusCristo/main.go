package main

import "fmt"

func main() {
	var n int
	var e float64 = 0

	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		factorial := 1
		// Se alcula el factorial
		for a := 1; a <= i; a++ {
			factorial *= a
		}
		e += (1 / float64(factorial))
	}
	fmt.Println(e)
}
