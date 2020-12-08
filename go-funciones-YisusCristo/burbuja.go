package main

import "fmt"

func Burbuja(s []int64) {
	n := len(s)
	sorted := false

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if s[i] > s[i+1] {
				s[i+1], s[i] = s[i], s[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

func main() {
	slice := []int64{1, 3, 69, -420, -310, -10, 200, 180, 95}
	// fmt.Println(slice)
	Burbuja(slice)
	fmt.Println(slice)
}
