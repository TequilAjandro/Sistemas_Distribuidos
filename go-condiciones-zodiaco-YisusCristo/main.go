package main

import "fmt"

func main() {
	var day int
	var month int

	fmt.Scan(&day)
	fmt.Scan(&month)
	if (month == 3 && day >= 21) || (month == 4 && day <= 20) {
		fmt.Println("aries")
	} else if (month == 4 && day >= 21) || (month == 5 && day <= 20) {
		fmt.Println("tauro")
	} else if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
		fmt.Println("geminis")
	} else if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		fmt.Println("cancer")
	} else if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		fmt.Println("leo")
	} else if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		fmt.Println("virgo")
	} else if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		fmt.Println("libra")
	} else if (month == 10 && day >= 23) || (month == 11 && day <= 22) {
		fmt.Println("escorpio")
	} else if (month == 11 && day >= 23) || (month == 12 && day <= 21) {
		fmt.Println("sagitario")
	} else if (month == 12 && day >= 22) || (month == 1 && day <= 20) {
		fmt.Println("sapricornio")
	} else if (month == 1 && day >= 21) || (month == 2 && day <= 18) {
		fmt.Println("acuario")
	} else {
		fmt.Println("piscis")
	}
}
