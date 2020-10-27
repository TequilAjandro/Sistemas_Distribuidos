package main

import "fmt"

func fibonacci(n int) int64 {
	if n == 0 || n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// Calcula los primeros 43 números
	for i := 0; i < 43; i++ {
		fib_number := fibonacci(i)
		fmt.Println(i+1, " th -> ", fib_number)
	}
}
