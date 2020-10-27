package main

import "fmt"

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	a := 123
	b := 456
	swap(&a, &b)
	fmt.Println("a -> ", a, "  b -> ", b)
}
