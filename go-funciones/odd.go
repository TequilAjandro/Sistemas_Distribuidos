package main

import (
	"fmt"
	"math/rand"
)

// Recibe un número el cual lo convertira a impar
func ToOdd(number int) int {
	oddNum := (number * 2) + 1
	return oddNum
}

func main() {
	min := 0
	max := 9999
	// Se generan 20 numeros aleatorios dentro de un rango de valores
	// y se convierte a impar
	for i := 0; i < 20; i++ {
		number := rand.Intn(max-min+1) + min
		oddNum := ToOdd(number)
		fmt.Println(oddNum)
	}
}
