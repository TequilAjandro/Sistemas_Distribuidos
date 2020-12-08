package main

import "fmt"

func main() {
	var base uint64
	var altura uint64

	fmt.Scanln(&base)
	fmt.Scanln(&altura)

	area := (base * altura) / 2
	fmt.Println(area)
}
