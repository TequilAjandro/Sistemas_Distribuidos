package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"

	process "./packages"
)

var processes []process.Process
var status bool = false

func creaLista() []process.Process {
	list := make([]process.Process, 0)
	for i := 0; i < 6; i++ {
		p := process.Process{Id: int(i), Value: int(0)}
		list = append(list, p)
	}
	return list
}

func arrancaprocesss() {
	for {
		for i, p := range processes {
			fmt.Println(p.Id, " : ", p.Value)
			processes[i].Value++
		}
		fmt.Println("-------------------------------------------------")
		time.Sleep(time.Millisecond * 500)
	}
}

func server() {
	s, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	processes = creaLista()
	go arrancaprocesss()

	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {

	if !status {
		err := gob.NewDecoder(c).Decode(&status)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if status {
		var p process.Process
		err := gob.NewDecoder(c).Decode(&p)
		if err != nil {
			fmt.Println(err)
		}
		processes = append(processes, p)
		status = false

	} else {
		proc := processes[len(processes)-1]
		err := gob.NewEncoder(c).Encode(proc)
		if err != nil {
			fmt.Println(err)
		} else {
			if len(processes) > 0 {
				processes = processes[:len(processes)-1]
			}
		}
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
	fmt.Println("Ended Server\n")
}
