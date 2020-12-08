package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var show bool = false

type Process struct {
	id    int
	value int
	quit  chan bool
}

func (p *Process) Run() {
	p.quit = make(chan bool)
	go func() {
		for {
			// fmt.Printf("id %d: %d", p.id, p.value)
			select {
			case <-p.quit:
				return
			default:
				if show {
					p.Print()
				}
				p.value++
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
}

func (p *Process) Stop() {
	p.quit <- true
}

func (p *Process) Print() {
	fmt.Printf("id %d: %d\n", p.id, p.value)
}

func delete(processes []Process, id int) {
	if len(processes) == 0 {
		fmt.Println("No hay ningun proceso corriendo")
	} else {
		for i, p := range processes {
			if p.id == id {
				p.Stop()
				processes = append(processes[:i], processes[i+1:]...)
			}
		}
	}
}

func menu() {
	fmt.Println("1.- Agregar proceso\n2.- Mostrar procesos")
	fmt.Println("3.- Terminar proceso\n4.- Salir")
}

func main() {
	idCount := 0
	p := []Process{}
	var channel chan bool

	menu()
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		fmt.Println("\n\n")
		menu()
		switch {
		case input.Text() == "1":
			fmt.Println("Agregar proceso")
			p = append(p, Process{len(p), 0, channel})
			p[idCount].Run()
			// go p[len(p)-1].Run()
			idCount++
		case input.Text() == "2":
			fmt.Println("Mostrar procesos")
			if idCount > 0 {
				show = true
			} else {
				fmt.Println("No hay ningun proceso corriendo")
			}
		case input.Text() == "3":
			fmt.Println("Eliminar proceso")
			if idCount > 0 {
				fmt.Print("ID: ")
				var id int
				fmt.Scanln(&id)
				delete(p, id)
			} else {
				fmt.Println("No hay ningun proceso corriendo")
			}
		case input.Text() == "q":
			show = false
		case input.Text() == "4":
			os.Exit(2)
		default:
			fmt.Println("Opción inexistente")
		}
	}
}
