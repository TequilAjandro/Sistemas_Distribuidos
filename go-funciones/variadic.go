package main

import "fmt"

func biggestNum(args ...int) int {
	biggest := 0
	for _, num := range args {
		if num > biggest {
			biggest = num
		}
	}
	return biggest
}

func main() {
	a := []int{1, 4, 2}
	fmt.Println(biggestNum(a...))
}
