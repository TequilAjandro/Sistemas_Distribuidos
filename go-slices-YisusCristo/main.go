package main

import "fmt"

func main() {
	var num int
	// Input del tamaño de s
	var n int
	fmt.Scanln(&n)
	// Creación del slice
	s := make([]int, n)
	// Llenar s con inputs del usuario
	for i := 0; i < n; i++ {
		fmt.Scanln(&num)
		s[i] = num
	}
	// Sumar
	var sum int
	for i := 0; i < n; i++ {
		sum += s[i]
	}
	fmt.Println(sum)
}
