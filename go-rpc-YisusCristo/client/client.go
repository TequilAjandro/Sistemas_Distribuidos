package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
)

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =", result)
	}

	// Menu
	fmt.Println()
	fmt.Println("1.- Agregar calificación de una materia\n2.- Mostrar el promedio del alumno")
	fmt.Println("3.- Mostrar el promedio general de un alumno\n4.- Mostrar promedio de una materia")
	fmt.Println("5.- Salir")

	input := bufio.NewScanner(os.Stdin)
	var res string

	for input.Scan() {
		switch {
		case input.Text() == "1":
			fmt.Println("Opción inexistente")
			c.Call("Server.Anos", int64(999), &res)
			fmt.Println(res)
		case input.Text() == "2":
			fmt.Println("Opción inexistente")
		case input.Text() == "3":
			fmt.Println("Opción inexistente")
		case input.Text() == "4":
			fmt.Println("Opción inexistente")
		case input.Text() == "5":
			fmt.Println("Opción inexistente")
		default:
			fmt.Println("Opción inexistente")
		}
	}
}

func main() {
	// go server()
	// go client()
	client()
	//
	// var input string
	// fmt.Scanln(&input)
}
